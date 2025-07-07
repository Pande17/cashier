// Fetch invoices when the page loads
window.onload = fetchInvoices;

// Function to handle view button click (example)
function viewInvoiceDetails(kodeInvoice) {
	alert(`Viewing details for invoice: ${kodeInvoice}`);
	// Add your code to fetch and display invoice details
}
let invoiceData = []; // Store all invoice data globally

// function to handle fetching invoices
async function fetchInvoices() {
	const apiUrl = 'http://localhost:3000/api/invoice'; // Ganti dengan endpoint API kamu
	const response = await fetch(apiUrl);
	const data = await response.json();
	invoiceData = data.data; // Store fetched data in global variable

	const invoiceGrid = document.getElementById('invoiceGrid');
	const invoiceCount = document.getElementById('invoiceCount');

	// Update the invoice count
	invoiceCount.textContent = invoiceData.length;

	invoiceGrid.innerHTML = ''; // Clear the grid before adding new items

	// Loop for dynamically generating invoice items
	invoiceData.forEach(invoice => {
		const formattedDate = new Date(invoice.tanggal_beli).toLocaleDateString('id-ID', {
			year: 'numeric',
			month: 'long',
			day: 'numeric',
		});

		// Update bagian ini di dalam loop yang menampilkan data invoice
		const invoiceItem = document.createElement('div');
		invoiceItem.classList.add('invoice-item', 'bg-white', 'rounded-lg', 'shadow-md', 'p-4', 'mb-4', 'hover:shadow-lg', 'transition-shadow', 'duration-300');
		invoiceItem.innerHTML = `
    <h2 class="text-lg font-semibold mb-2">${invoice.member_id || 'Umum'}</h2>
    <p class="text-gray-500">${invoice.kode_invoice}</p>
    <p class="text-gray-600">${formattedDate}</p>
    <p class="invoice-total mb-4">Rp. ${parseInt(invoice.total).toLocaleString('id-ID')}</p>
    <div class="invoice-actions">
        <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" onclick="viewInvoice('${invoice.kode_invoice}')"><i class="ri-eye-line"></i> </button>
        <button class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded" onclick="updateInvoice('${invoice.kode_invoice}')"><i class="ri-pencil-line"></i> </button>
        <button class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded" onclick="deleteInvoice('${invoice.kode_invoice}')"><i class="ri-delete-bin-line"></i> </button>
    </div>
`;
		invoiceGrid.appendChild(invoiceItem);
	});
}

// Search and Filter Function
function filterInvoices() {
	const searchInput = document.getElementById('searchInput').value.toLowerCase();
	const filteredInvoices = invoiceData.filter(invoice => {
		return invoice.kode_invoice.toLowerCase().includes(searchInput) || invoice.member_id.toLowerCase().includes(searchInput);
	});

	// Update the invoice count with filtered results
	const invoiceCount = document.getElementById('invoiceCount');
	invoiceCount.textContent = filteredInvoices.length;

	// Rebuild the grid with filtered invoices
	const invoiceGrid = document.getElementById('invoiceGrid');
	invoiceGrid.innerHTML = ''; // Clear the grid before adding new items

	filteredInvoices.forEach(invoice => {
		const formattedDate = new Date(invoice.tanggal_beli).toLocaleDateString('id-ID', {
			year: 'numeric',
			month: 'long',
			day: 'numeric',
		});

		const invoiceItem = document.createElement('div');
		invoiceItem.classList.add('invoice-item', 'bg-white', 'rounded-lg', 'shadow-md', 'p-4', 'mb-4', 'hover:shadow-lg', 'transition-shadow', 'duration-300');
		invoiceItem.innerHTML = `
            <h2 class="text-lg font-semibold mb-2">${invoice.member_id || 'Umum'}</h2>
            <p class="text-gray-600">${formattedDate}</p>
            <div class="invoice-header">
                <p class="invoice-total">Rp. ${parseInt(invoice.total).toLocaleString('id-ID')}</p>
                <div class="invoice-actions">
                    <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" onclick="viewInvoice(${invoice.kode_invoice})"><i class="ri-eye-line"></i></button>
                    <button class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded" onclick="updateInvoice(${invoice.kode_invoice})"><i class="ri-pencil-line"></i></button>
                    <button class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded" onclick="deleteInvoice('${invoice.kode_invoice}')"><i class="ri-delete-bin-line"></i></button>
                </div>
            </div>
            <p class="text-gray-500">${invoice.kode_invoice}</p>
        `;
		invoiceGrid.appendChild(invoiceItem);
	});
}

// Fungsi untuk menangani klik tombol View
async function viewInvoice(kodeInvoice) {
	// Mengambil data detail invoice berdasarkan kode_invoice
	try {
		const response = await fetch(`http://localhost:3000/api/invoice/${kodeInvoice}`);
		const data = await response.json();

		if (data && data.data) {
			// Ambil data dari respons API
			const invoice = data.data;

			// Tampilkan konten dalam modal
			const invoiceDetailContent = document.getElementById('invoiceDetailContent');
			invoiceDetailContent.innerHTML = `
                <p><strong>Kode Invoice:</strong> ${invoice.kode_invoice}</p>
                <p><strong>Member ID:</strong> ${invoice.member_id || 'Umum'}</p>
                <p><strong>Tanggal Beli:</strong> ${new Date(invoice.tanggal_beli).toLocaleDateString('id-ID')}</p>
                <p><strong>Total:</strong> Rp. ${parseInt(invoice.total).toLocaleString('id-ID')}</p>
                <p><strong>Keterangan:</strong> ${invoice.keterangan || 'Tidak ada keterangan'}</p>
            `;

			// Menampilkan modal
			document.getElementById('invoiceDetailModal').classList.remove('hidden');
		} else {
			alert('Invoice not found');
		}
	} catch (error) {
		console.error('Error fetching invoice details:', error);
		alert('Error fetching invoice details.');
	}
}

// Menangani klik tombol Close pada modal
document.getElementById('closeModal').addEventListener('click', function () {
	// Tutup modal
	document.getElementById('invoiceDetailModal').classList.add('hidden');
});

// Session for handling invoice deletion
let invoiceToDelete = null; // Menyimpan kode invoice yang akan dihapus

// Fungsi untuk menangani klik tombol Delete
function deleteInvoice(kodeInvoice) {
	invoiceToDelete = kodeInvoice; // Menyimpan kode invoice yang akan dihapus
	// Menampilkan modal konfirmasi
	document.getElementById('deleteModal').classList.add('active'); // Menampilkan modal
}

// Menangani klik tombol "Delete" pada modal konfirmasi
document.getElementById('confirmDelete').addEventListener('click', async function () {
	if (invoiceToDelete) {
		try {
			const response = await fetch(`http://localhost:3000/api/invoice/${invoiceToDelete}`, {
				method: 'DELETE', // Menggunakan metode DELETE untuk menghapus data
			});

			if (response.ok) {
				alert(`Invoice ${invoiceToDelete} deleted successfully.`);
				document.getElementById('deleteModal').classList.remove('active'); // Menutup modal
				fetchInvoices(); // Menarik ulang data setelah penghapusan
			} else {
				alert(`Failed to delete invoice ${invoiceToDelete}.`);
			}
		} catch (error) {
			alert('Error deleting invoice.');
		}
	}
});

// Menangani klik tombol "Cancel" pada modal konfirmasi
document.getElementById('cancelDelete').addEventListener('click', function () {
	// Tutup modal jika cancel diklik
	document.getElementById('deleteModal').classList.remove('active'); // Menutup modal
});
