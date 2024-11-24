# Base image
FROM node:18 AS builder

# Set working directory
WORKDIR /app

# Copy package.json and install dependencies
COPY package.json package-lock.json ./
RUN npm install

# Copy the source code
COPY . .

# Build the Svelte application
RUN npm run build

# Final image
FROM nginx:stable-alpine

# Copy built files to NGINX html folder
COPY --from=builder /app/public /usr/share/nginx/html

# Expose the frontend port
EXPOSE 80

# Run NGINX
CMD ["nginx", "-g", "daemon off;"]
