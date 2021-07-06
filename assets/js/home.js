const toggleBtn = document.querySelector('.nav_toggleBtn');
const menu = document.querySelector('.nav_menu');
const icon = document.querySelector('.nav_icon');
const wrap = document.querySelector('.wrap');

toggleBtn.addEventListener('click',()=>{
    menu.classList.toggle('active');
    icon.classList.toggle('active');
    wrap.classList.toggle('active');
});