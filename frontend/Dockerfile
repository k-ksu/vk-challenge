# Use Node.js as the base image since http-server is a Node.js package
FROM node:16-alpine

# Create a directory for the app
WORKDIR /app

# Copy the index.html file into the container
COPY index.html /app/

# Install http-server globally
RUN npm install -g http-server

# Expose port 8080 for the web server
EXPOSE 8080

# Start the http-server when the container starts
CMD ["http-server", "-p", "8080"]
