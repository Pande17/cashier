# Gunakan image nginx sebagai base image
FROM nginx:alpine

# Salin file frontend (HTML, CSS, JS) ke dalam kontainer
COPY . /usr/share/nginx/html

# Salin file konfigurasi nginx ke dalam kontainer
COPY nginx.conf /etc/nginx/nginx.conf

# Expose port 80 untuk akses ke web
EXPOSE 80

# Jalankan nginx di background
CMD ["nginx", "-g", "daemon off;"]