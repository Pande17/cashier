# Stage 1 - Build stage
FROM node:18-alpine AS build

# Set working directory
WORKDIR /app

# Install dependencies
COPY package.json ./ 
COPY package-lock.json ./   
#It's a good practice to copy the lock file as well

RUN npm install

# Copy all source code to build it
COPY . ./ 

# Run Tailwind CSS in watch mode (for development)
CMD ["npx", "tailwindcss", "-i", "./src/input.css", "-o", "./src/output.css", "--watch"]

# Stage 2 - Production stage (optional)
FROM nginx:alpine AS production

# Copy built files from build stage to nginx's serving folder
COPY --from=build /app/src/output.css /usr/share/nginx/html/output.css
# Adjust if necessary
COPY --from=build /app ./usr/share/nginx/html  

# Expose port
EXPOSE 80

# Start nginx to serve static files
CMD ["nginx", "-g", "daemon off;"]