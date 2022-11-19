-- Setup a dev server --

TODO -- DEPRECATED

0. Tools needed : docker, go, prisma (npm install prisma --global)

1. a) Create a spotify app to obtain API credentials (https://developer.spotify.com/dashboard/)
   b) Add https://localhost:8080/callback/spotify to the list of Redirect URIs in the app settings

2. Run `cp default_config.yml config.yml` from inside the backend folder and replace the Client ID and Client Secret with the correct values

3. Create a self-signed SSL certificate and set the correct paths to the cert in config.yml if needed (by default it will look for cert.pem and key.pem in the backend folder)

4. a) Launch the Prisma and Postgres containers as daemons with `docker compose up -d`
   b) Deploy the Prisma service with `prisma deploy` - Prisma should be running on port 4466 (admin console at `http://localhost:4466/_admin`)

5. Build the server with `go build` (or run directly with `go run main.go`)
