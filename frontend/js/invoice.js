// Membuat elemen <style> untuk CSS
const style = document.createElement('style');
style.innerHTML = `

    .btn-delete {
        background-color: red;
        color: white;
        padding: 5px 10px;
        border-radius: 5px;
    }

	.btn-delete:hover {
		background-color: darkred;
	}

`;

// Menambahkan elemen <style> ke dalam <head> HTML
document.head.appendChild(style);

// Function to fetch and display invoices
async function fetchInvoices() {
	const apiUrl = 'http://localhost:3000/api/invoice'; // Replace with your API endpoint

	const response = await fetch(apiUrl);
	const data = await response.json();

	const tableBody = document.getElementById('invoiceTableBody');
	tableBody.innerHTML = ''; // Clear the table before adding new rows

	// Loop through the fetched data and create rows dynamically
	data.data.forEach((invoices, index) => {
		const row = document.createElement('tr');

		// Assign background color based on row index (odd or even)
		const rowClass = index % 2 === 0 ? 'table-row-even' : 'table-row-odd';

		row.classList.add(rowClass);

		// Add content to the row
		row.innerHTML = `
            <td class="px-6 py-4 border-b border-gray-300 text-lg">${invoices.kode_invoice}</td>
            <td class="px-6 py-4 border-b border-gray-300 text-lg">${invoices.member_id}</td>
            <td class="px-6 py-4 border-b border-gray-300 text-lg">${new Date(invoices.tanggal_beli).toLocaleDateString()}</td>
            <td class="px-6 py-4 border-b border-gray-300 text-lg">${new Date(invoices.jatuh_tempo).toLocaleDateString()}</td>
            <td class="px-6 py-4 border-b border-gray-300 text-lg">${invoices.status || 'N/A'}</td>
            <td class="px-6 py-4 border-b border-gray-300 text-lg">${invoices.total.toLocaleString()}</td>
            <td class="px-6 py-4 border-b border-gray-300 text-lg">
                <button class="btn btn-primary btn-sm" onclick="viewInvoiceDetails('${invoices.kode_invoice}')">View</button>
                <button class="btn btn-delete btn-sm" onclick="viewInvoiceDetails('${invoices.kode_invoice}')">Delete</button>
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
