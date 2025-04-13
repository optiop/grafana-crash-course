import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./App";
import {
  getWebInstrumentations,
  initializeFaro,
  ReactIntegration,
  ReactRouterVersion,
} from "@grafana/faro-react";
import { TracingInstrumentation } from "@grafana/faro-web-tracing";
import {
  createRoutesFromChildren,
  matchRoutes,
  Routes,
  useLocation,
  useNavigationType,
} from "react-router-dom";

initializeFaro({
  // url: "http://localhost:12345/alloy/api/v1/push",
  url: "http://localhost:4319/collect",
  app: {
    name: "fo11y-demo",
    version: "1.0.0",
    environment: "production",
  },

  instrumentations: [
    ...getWebInstrumentations(),

    new TracingInstrumentation(),

    new ReactIntegration({
      router: {
        version: ReactRouterVersion.V6,
        dependencies: {
          createRoutesFromChildren,
          matchRoutes,
          Routes,
          useLocation,
          useNavigationType,
        },
      },
    }),
  ],
});

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
