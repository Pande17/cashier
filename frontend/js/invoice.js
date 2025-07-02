// Function to fetch and display invoices
async function fetchInvoices() {
	const apiUrl = 'http://localhost:3000/api/invoice'; // Replace with your API endpoint

	const response = await fetch(apiUrl);
	const data = await response.json();

	const tableBody = document.getElementById('invoiceTableBody');
	tableBody.innerHTML = ''; // Clear the table before adding new rows

	data.data.forEach(invoices => {
		const row = document.createElement('tr');

		row.innerHTML = `
                    <td class="px-6 py-4 border-b border-gray-300 text-lg">${invoices.kode_invoice}</td>
                    <td class="px-6 py-4 border-b border-gray-300 text-lg">${invoices.member_id}</td>
                    <td class="px-6 py-4 border-b border-gray-300 text-lg">${new Date(invoices.tanggal_beli).toLocaleDateString()}</td>
                    <td class="px-6 py-4 border-b border-gray-300 text-lg">${new Date(invoices.jatuh_tempo).toLocaleDateString()}</td>
                    <td class="px-6 py-4 border-b border-gray-300 text-lg">${invoices.status || 'N/A'}</td>
                    <td class="px-6 py-4 border-b border-gray-300 text-lg">${invoices.total.toLocaleString()}</td>
                    <td class="px-6 py-4 border-b border-gray-300 text-lg">
                        <button class="btn btn-primary btn-sm" onclick="viewInvoiceDetails('${invoices.kode_invoice}')">View</button>
                    </td>
                `;

		tableBody.appendChild(row);
	});
}


// Fetch invoices when the page loads
window.onload = fetchInvoices;

// Function to handle view button click (example)
function viewInvoiceDetails(kodeInvoice) {
	alert(`Viewing details for invoice: ${kodeInvoice}`);
	// Add your code to fetch and display invoice details
}
