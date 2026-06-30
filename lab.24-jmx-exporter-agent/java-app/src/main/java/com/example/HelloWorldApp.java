package com.example;

import com.sun.net.httpserver.HttpServer;
import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;
import java.util.concurrent.atomic.AtomicLong;
import java.lang.management.ManagementFactory;
import javax.management.MBeanServer;
import javax.management.ObjectName;

public class HelloWorldApp implements HelloWorldAppMBean {

    private static final AtomicLong requestCount = new AtomicLong(0);
    private static final AtomicLong totalResponseTimeMs = new AtomicLong(0);

    @Override
    public long getTotalRequests() {
        return requestCount.get();
    }

    @Override
    public long getTotalResponseTimeMs() {
        return totalResponseTimeMs.get();
    }

    public static void main(String[] args) throws Exception {
        // Register custom MBean so JMX Exporter can expose it
        HelloWorldApp app = new HelloWorldApp();
        MBeanServer mbs = ManagementFactory.getPlatformMBeanServer();
        ObjectName name = new ObjectName("com.example:type=RequestStats");
        mbs.registerMBean(app, name);

        HttpServer server = HttpServer.create(new InetSocketAddress(8080), 0);

        server.createContext("/", exchange -> {
            long start = System.currentTimeMillis();
            long count = requestCount.incrementAndGet();

            String response = String.format(
                "Hello World!\n" +
                "------------------------------\n" +
                "Request count : %d\n" +
                "JVM metrics   : http://localhost:9090/metrics\n" +
                "Prometheus    : http://localhost:9091\n",
                count
            );

            byte[] bytes = response.getBytes();
            exchange.sendResponseHeaders(200, bytes.length);
            try (OutputStream os = exchange.getResponseBody()) {
                os.write(bytes);
            }
            totalResponseTimeMs.addAndGet(System.currentTimeMillis() - start);
        });

        server.createContext("/health", exchange -> {
            String response = "{\"status\":\"UP\"}\n";
            exchange.getResponseHeaders().set("Content-Type", "application/json");
            byte[] bytes = response.getBytes();
            exchange.sendResponseHeaders(200, bytes.length);
            try (OutputStream os = exchange.getResponseBody()) {
                os.write(bytes);
            }
        });

        server.start();
        System.out.println("==============================================");
        System.out.println("  Hello World JMX Exporter Demo");
        System.out.println("==============================================");
        System.out.println("  App     -> http://localhost:8080");
        System.out.println("  Metrics -> http://localhost:9090/metrics");
        System.out.println("  Prometheus -> http://localhost:9091");
        System.out.println("==============================================");
    }
}
