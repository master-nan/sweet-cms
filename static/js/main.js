(function () {
    const sidebarNavWrapper = document.querySelector(".sidebar-nav-wrapper");
    const mainWrapper = document.querySelector(".main-wrapper");
    const menuToggleButton = document.querySelector("#menu-toggle");
    const menuToggleButtonIcon = document.querySelector("#menu-toggle i");
    const overlay = document.querySelector(".overlay");

    menuToggleButton.addEventListener("click", () => {
        sidebarNavWrapper.classList.toggle("active");
        overlay.classList.add("active");
        mainWrapper.classList.toggle("active");

        if (document.body.clientWidth > 1200) {
            if (menuToggleButtonIcon.classList.contains("bi-list")) {
                menuToggleButtonIcon.classList.remove("bi-list");
                menuToggleButtonIcon.classList.add("bi-chevron-right");
            } else {
                menuToggleButtonIcon.classList.remove("bi-chevron-right");
                menuToggleButtonIcon.classList.add("bi-list");
            }
        } else {
            if (menuToggleButtonIcon.classList.contains("bi-list")) {
                menuToggleButtonIcon.classList.remove("bi-chevron-right");
                menuToggleButtonIcon.classList.add("bi-list");
            }
        }
    });
    overlay.addEventListener("click", () => {
        sidebarNavWrapper.classList.remove("active");
        overlay.classList.remove("active");
        mainWrapper.classList.remove("active");
    });
})()