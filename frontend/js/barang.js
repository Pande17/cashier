// Fetch barang when the page loads
window.onload = fetchBarang;

let barangData = []; // Store all barang data globally

// Function to handle fetching barang
async function fetchBarang() {
	const apiUrl = 'http://localhost:3000/api/barang'; // Ganti dengan endpoint API kamu
	const response = await fetch(apiUrl);
	const data = await response.json();
	barangData = data.data; // Store fetched data in global variable

	const barangGrid = document.getElementById('barangGrid');
	const barangCount = document.getElementById('barangCount');

	// Update the barang count
	barangCount.textContent = barangData.length;

	barangGrid.innerHTML = ''; // Clear the grid before adding new items

	// Loop for dynamically generating barang items
	barangData.forEach(barang => {
		// Tampilkan data yang diperlukan: Nama, Kode, Harga Jual, Kategori
		const barangItem = document.createElement('div');
		barangItem.classList.add('barang-item', 'bg-white', 'rounded-lg', 'shadow-md', 'p-4', 'mb-4', 'hover:shadow-lg', 'transition-shadow', 'duration-300');
		barangItem.innerHTML = `
            <h2 class="text-lg font-semibold mb-2">${barang.nama || 'Tidak Diketahui'}</h2>
            <p class="text-gray-600">Kode Barang: ${barang.kode_barang}</p>
            <p class="text-gray-600">Kategori: ${barang.kategori}</p>
            <p class="text-lg font-bold text-red-600">Rp. ${parseInt(barang.harga_jual).toLocaleString('id-ID')}</p>
            <div class="barang-actions mt-4">
                <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" onclick="viewBarang(${barang.kode_barang})">View</button>
                <button class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded" onclick="updateBarang(${barang.kode_barang})">Update</button>
                <button class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded" onclick="deleteBarang(${barang.kode_barang})">Delete</button>
            </div>
        `;
		barangGrid.appendChild(barangItem);
	});
}

// Search and Filter Function
function filterBarang() {
	const searchInput = document.getElementById('searchInput').value.toLowerCase();
	const filteredBarang = barangData.filter(barang => {
		return barang.kode_barang.toLowerCase().includes(searchInput) || barang.nama.toLowerCase().includes(searchInput);
	});

	// Update the barang count with filtered results
	const barangCount = document.getElementById('barangCount');
	barangCount.textContent = filteredBarang.length;

	// Rebuild the grid with filtered barang
	const barangGrid = document.getElementById('barangGrid');
	barangGrid.innerHTML = ''; // Clear the grid before adding new items

	filteredBarang.forEach(barang => {
		// Tampilkan data yang diperlukan: Nama, Kode, Harga Jual, Kategori
		const barangItem = document.createElement('div');
		barangItem.classList.add('barang-item', 'bg-white', 'rounded-lg', 'shadow-md', 'p-4', 'mb-4', 'hover:shadow-lg', 'transition-shadow', 'duration-300');
		barangItem.innerHTML = `
            <h2 class="text-lg font-semibold mb-2">${barang.nama || 'Tidak Diketahui'}</h2>
            <p class="text-gray-600">Kode Barang: ${barang.kode_barang}</p>
            <p class="text-gray-600">Kategori: ${barang.kategori}</p>
            <p class="text-lg font-bold text-red-600">Rp. ${parseInt(barang.harga_jual).toLocaleString('id-ID')}</p>
            <div class="barang-actions mt-4">
                <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" onclick="viewBarang(${barang.kode_barang})">View</button>
                <button class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded" onclick="updateBarang(${barang.kode_barang})">Update</button>
                <button class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded" onclick="deleteBarang(${barang.kode_barang})">Delete</button>
            </div>
        `;
		barangGrid.appendChild(barangItem);
	});
}

// Session for handling barang deletion
let barangToDelete = null; // Menyimpan kode barang yang akan dihapus

// Fungsi untuk menangani klik tombol Delete
function deleteBarang(kodeBarang) {
	barangToDelete = kodeBarang; // Menyimpan kode barang yang akan dihapus
	// Menampilkan modal konfirmasi
	document.getElementById('deleteModal').classList.remove('hidden');
}

// Menangani klik tombol "Delete" pada modal konfirmasi
document.getElementById('confirmDelete').addEventListener('click', async function () {
	if (barangToDelete) {
		try {
			// Panggil API DELETE untuk soft delete barang
			const response = await fetch(`http://localhost:3000/api/barang/${barangToDelete}`, {
				method: 'DELETE', // Menggunakan metode DELETE untuk menghapus data
			});

			if (response.ok) {
				// Jika berhasil menghapus
				alert(`Barang ${barangToDelete} deleted successfully.`);
				// Tutup modal dan lakukan pengambilan ulang data
				document.getElementById('deleteModal').classList.add('hidden');
				fetchBarang(); // Menarik ulang data setelah penghapusan
			} else {
				// Jika gagal menghapus
				alert(`Failed to delete barang ${barangToDelete}.`);
			}
		} catch (error) {
			alert('Error deleting barang.');
		}
	}
});

// Menangani klik tombol "Cancel" pada modal konfirmasi
document.getElementById('cancelDelete').addEventListener('click', function () {
	// Tutup modal jika cancel diklik
	document.getElementById('deleteModal').classList.add('hidden');
});
