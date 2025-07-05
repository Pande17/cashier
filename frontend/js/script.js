// Fetch and load the sidebar template
fetch('../templates/sidebar.html')
	.then(response => response.text())
	.then(data => {
		document.getElementById('sidebarContainer').innerHTML = data;
	})
	.catch(error => console.error('Error loading sidebar:', error));

// // Sidebar toggle functionality
// function toggleSidebar() {
// 	const sidebar = document.querySelector('.sidebar');
// 	sidebar.classList.toggle('w-[65px]'); // Collapse the sidebar by adjusting width
// 	sidebar.classList.toggle('w-[300px]'); // Expand the sidebar when toggled
// 	const arrow = document.querySelector('#arrow');
// 	arrow.classList.toggle('rotate-180'); // Rotate arrow icon when sidebar is collapsed
// }

function toggleDropdown() {
	document.querySelector('#submenu').classList.toggle('hidden');
	document.querySelector('#arrow').classList.toggle('rotate-0');
}
