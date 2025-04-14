## 🚀 Getting Started

1. **Start the Docker Stack**

   Navigate to the `solution` directory and run:

   ```bash
   docker compose up -d
   ```

   This command will start both Grafana and Keycloak containers.

2. **Access the Keycloak Admin Console**

   Open your browser and go to [http://localhost:8080](http://localhost:8080).

   - Log in with the default administrator credentials
     - Username: `admin`
     - Password: `admin`

3. **Create a New User in Keycloak**

   In the Keycloak admin console:

   - Navigate to **Users** > **Add User**
   - Fill in the required fields
   - **Important**: Assign a valid email address to the user
   - Click **Save**

   After saving, go to the **Credentials** tab:

   - Set a password for the user
   - Toggle **Temporary** to **Off** to prevent the user from being prompted to change the password on first login
   - Click **Set Password**

4. **Access Grafana**

   Open your browser and go to [http://localhost:3000](http://localhost:3000).

   - Click on **Sign in with Keycloak**
   - You will be redirected to the Keycloak login page
   - Log in with the user credentials you created earlier
   - Upon successful authentication, you will be redirected back to Grafana and logged in

## 📝 Notes

 The Keycloak realm `grafana` is pre-provisioned; there's no need to create it manuall.
 No modifications to `/etc/hosts` are necessary; services are accessible via `localhost.
