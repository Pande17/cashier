// Function to fetch and display members
async function fetchmembers() {
	const apiUrl = 'http://localhost:3000/api/member'; // Replace with your API endpoint

	const response = await fetch(apiUrl);
	const data = await response.json();

	const tableBody = document.getElementById('memberTableBody');
	tableBody.innerHTML = ''; // Clear the table before adding new rows

	data.data.forEach(members => {
		const row = document.createElement('tr');

		row.innerHTML = `
                    <td class="px-6 py-4 border-b border-gray-300 text-lg">${members.id}</td>
                    <td class="px-6 py-4 border-b border-gray-300 text-lg">${members.nama}</td>
                    <td class="px-6 py-4 border-b border-gray-300 text-lg">${members.pic}</td>
                    <td class="px-6 py-4 border-b border-gray-300 text-lg">${members.perusahaan}</td>
                    <td class="px-6 py-4 border-b border-gray-300 text-lg">${members.kategori}</td>
                    <td class="px-6 py-4 border-b border-gray-300 text-lg">${members.alamat}</td>
                    <td class="px-6 py-4 border-b border-gray-300 text-lg">${members.no_telp}</td>
                    <td class="px-6 py-4 border-b border-gray-300 text-lg">${members.status || 'N/A'}</td>
                    <td class="px-6 py-4 border-b border-gray-300 text-lg">
                        <button class="btn btn-primary btn-sm" onclick="viewMemberDetails('${members.id}')">View</button>
                    </td>
                `;

		tableBody.appendChild(row);
	});
}

// Fetch members when the page loads
window.onload = fetchmembers;

// Function to handle view button click (example)
function viewmemberDetails(id) {
	alert(`Viewing details for member: ${id}`);
	// Add your code to fetch and display member details
}
