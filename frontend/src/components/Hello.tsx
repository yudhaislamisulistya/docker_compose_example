import React, { useEffect, useState } from "react";

export function Hello() {
  // State untuk menyimpan data
  const [students, setStudents] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    // Fungsi untuk mengambil data dari API
    const fetchStudents = async () => {
      try {
        const response = await fetch("http://localhost:1323/students");
        if (!response.ok) {
          throw new Error("Failed to fetch data");
        }
        const data = await response.json();
        setStudents(data); // Set data yang diterima ke state
      } catch (error) {
        setError(error.message); // Tangani error
      } finally {
        setLoading(false); // Set loading selesai
      }
    };

    fetchStudents().catch(console.error); // Tangani promise rejection
  }, []); // Hanya dijalankan sekali ketika komponen pertama kali dimounting

  // Jika loading, tampilkan pesan loading
  if (loading) {
    return <p>Loading...</p>;
  }

  // Jika ada error, tampilkan pesan error
  if (error) {
    return <p>Error: {error}</p>;
  }

  // Jika data sudah tersedia, tampilkan daftar students
  return (
    <div>
      <h1>Students List</h1>
      <ul>
        {students.map((student, index) => (
          <li key={index}>{student.name}</li> // Asumsi ada properti 'name' dalam data student
        ))}
      </ul>
    </div>
  );
}
