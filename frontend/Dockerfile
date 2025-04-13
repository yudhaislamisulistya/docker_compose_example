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

# Build aplikasi Next.js
RUN npm run build

# Jalankan aplikasi Next.js
CMD ["npm", "start"]
