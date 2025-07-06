// Fetch member data when the page loads
window.onload = fetchMember;

let memberData = []; // Store all member data globally

// Function to handle fetching member data
async function fetchMember() {
	const apiUrl = 'http://localhost:3000/api/member'; // Ganti dengan endpoint API kamu
	const response = await fetch(apiUrl);
	const data = await response.json();
	memberData = data.data; // Store fetched data in global variable

	const memberGrid = document.getElementById('memberGrid');
	const memberCount = document.getElementById('memberCount');

	// Update the member count
	memberCount.textContent = memberData.length;

	memberGrid.innerHTML = ''; // Clear the grid before adding new items

	// Loop for dynamically generating member items
	memberData.forEach(member => {
		const memberItem = document.createElement('div');
		memberItem.classList.add('member-item', 'bg-white', 'rounded-lg', 'shadow-md', 'p-4', 'mb-4', 'hover:shadow-lg', 'transition-shadow', 'duration-300');
		memberItem.innerHTML = `
            <h2 class="text-lg font-semibold mb-2">${member.nama || 'Tidak Diketahui'}</h2>
            <p class="text-gray-600">ID Member: ${member.id}</p>
            <p class="text-gray-600">Perusahaan: ${member.perusahaan}</p>
            <p class="text-gray-600">Kategori: ${member.kategori}</p>
            <div class="member-actions mt-4">
                <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" onclick="viewMember(${member.id})">View</button>
                <button class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded" onclick="updateMember(${member.id})">Update</button>
                <button class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded" onclick="deleteMember(${member.id})">Delete</button>
            </div>
        `;
		memberGrid.appendChild(memberItem);
	});
}

// Search and Filter Function
function filterMember() {
	const searchInput = document.getElementById('searchInput').value.toLowerCase();
	const filteredMember = memberData.filter(member => {
		return (
			member.id.toLowerCase().includes(searchInput) ||
			member.nama.toLowerCase().includes(searchInput) ||
			member.perusahaan.toLowerCase().includes(searchInput) ||
			member.kategori.toLowerCase().includes(searchInput)
		);
	});

	// Update the member count with filtered results
	const memberCount = document.getElementById('memberCount');
	memberCount.textContent = filteredMember.length;

	// Rebuild the grid with filtered members
	const memberGrid = document.getElementById('memberGrid');
	memberGrid.innerHTML = ''; // Clear the grid before adding new items

	filteredMember.forEach(member => {
		const memberItem = document.createElement('div');
		memberItem.classList.add('member-item', 'bg-white', 'rounded-lg', 'shadow-md', 'p-4', 'mb-4', 'hover:shadow-lg', 'transition-shadow', 'duration-300');
		memberItem.innerHTML = `
            <h2 class="text-lg font-semibold mb-2">${member.nama || 'Tidak Diketahui'}</h2>
            <p class="text-gray-600">ID Member: ${member.id}</p>
            <p class="text-gray-600">Perusahaan: ${member.perusahaan}</p>
            <p class="text-gray-600">Kategori: ${member.kategori}</p>
            <div class="member-actions mt-4">
                <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" onclick="viewMember(${member.id})">View</button>
                <button class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded" onclick="updateMember(${member.id})">Update</button>
                <button class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded" onclick="deleteMember(${member.id})">Delete</button>
            </div>
        `;
		memberGrid.appendChild(memberItem);
	});
}

// Session for handling member deletion
let memberToDelete = null; // Menyimpan id member yang akan dihapus

// Fungsi untuk menangani klik tombol Delete
function deleteMember(idMember) {
	memberToDelete = idMember; // Menyimpan id member yang akan dihapus
	// Menampilkan modal konfirmasi
	document.getElementById('deleteModal').classList.remove('hidden');
}

// Menangani klik tombol "Delete" pada modal konfirmasi
document.getElementById('confirmDelete').addEventListener('click', async function () {
	if (memberToDelete) {
		try {
			// Panggil API DELETE untuk soft delete member
			const response = await fetch(`http://localhost:3000/api/member/${memberToDelete}`, {
				method: 'DELETE', // Menggunakan metode DELETE untuk menghapus data
			});

			if (response.ok) {
				// Jika berhasil menghapus
				alert(`Member ${memberToDelete} deleted successfully.`);
				// Tutup modal dan lakukan pengambilan ulang data
				document.getElementById('deleteModal').classList.add('hidden');
				fetchMember(); // Menarik ulang data setelah penghapusan
			} else {
				// Jika gagal menghapus
				alert(`Failed to delete member ${memberToDelete}.`);
			}
		} catch (error) {
			alert('Error deleting member.');
		}
	}
});

// Menangani klik tombol "Cancel" pada modal konfirmasi
document.getElementById('cancelDelete').addEventListener('click', function () {
	// Tutup modal jika cancel diklik
	document.getElementById('deleteModal').classList.add('hidden');
});
