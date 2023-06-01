function selectFont(font) {
    var dropdownBtn = document.querySelector('.dropdown-btn');
    dropdownBtn.textContent = font;
    closeDropdown();
}

function closeDropdown() {
    var dropdownContent = document.querySelector('.dropdown-content');
    dropdownContent.style.display = 'none';
}

document.addEventListener('click', function (e) {
    var target = e.target;
    var dropdownContent = document.querySelector('.dropdown-content');

    if (!target.closest('.dropdown')) {
        dropdownContent.style.display = 'none';
    } else if (target.classList.contains('dropdown-btn')) {
        dropdownContent.style.display = (dropdownContent.style.display === 'block') ? 'none' : 'block';
    }
});
