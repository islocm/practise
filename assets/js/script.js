function opentable() {
    var open;
    open = document.getElementsByClassName("hide");

    document.getElementById("myBtn").style.width = "700%";

    for (let index = 0; index < open.length; index++) {
        open[index].style.display = "block";



    }
}