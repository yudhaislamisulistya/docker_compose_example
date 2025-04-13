# Gunakan image Node.js sebagai base
FROM node:18-alpine

# Set working directory di dalam container
WORKDIR /app

# Salin package.json dan package-lock.json
COPY package*.json ./

# Install dependensi
RUN npm install

# Salin seluruh kode aplikasi
COPY . .

# Set variabel lingkungan untuk Next.js (jika diperlukan)
ARG NEXT_PUBLIC_API_URL 
ENV NEXT_PUBLIC_API_URL=${NEXT_PUBLIC_API_URL}

# Build aplikasi Next.js
RUN npm run build

# Expose port yang digunakan oleh Next.js
EXPOSE 3000

# Jalankan aplikasi Next.js
CMD ["npm", "start"]
