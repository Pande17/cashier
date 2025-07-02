// Burger Menu Toggle
const burgerMenu = document.getElementById('burger-menu');
const sidebar = document.querySelector('.bg-gray-800');
burgerMenu.addEventListener('click', () => {
	sidebar.classList.toggle('hidden');
});
