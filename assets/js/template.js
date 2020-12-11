function templateopenbars() {
    let barsopen = document.getElementById("container-bar");
    barsopen.style.width = "25%";


}

function templatehidewindow() {
    let barsopen = document.getElementById("container-bar");
    barsopen.style.width = "0px";


}


function templateopenlog() {
    let open = document.getElementById("login-active");
    open.style.width = "125px";
}

function templatecloselog() {
    let open = document.getElementById("login-active");
    open.style.width = "0px";
}

function templateremovecookie() {
    let open = document.getElementById("login-active");
    open.style.width = "0px";
    let remove = new XMLHttpRequest();
    remove.open("POST", "/loginexit");

    remove.send()

    window.location.assign("/")
    window.location.assign("/")



}

function templatetosms() {
    window.location.assign("/sms");
}

function templatepreviousbutton() {
    history.back();
}

function templatenextbutton() {
    history.forward();
}

function downrightdrop() {
    let open = document.getElementById("menu-content");
    open.style.height = "26px";
    let del = document.getElementById("mahalla");
    del.remove();

}

document.getElementById("fadd").addEventListener("click", fadd);
document.getElementById("fupdate").addEventListener("click", fupdate);
document.getElementById("fdelete").addEventListener("click", fdelete);

function downrightclose() {
    let open = document.getElementById("menu-content");
    open.style.height = "0px";
}