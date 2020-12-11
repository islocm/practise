function windowqaror(qaror) {
    let creating = document.createElement("div");
    creating.classList.add("modelling");
    let cre = document.createElement("input");
    cre.className = "modalinput";
    let times = document.createElement("i");
    times.classList.add("fas", "fa-times");
    times.onclick = function() {
        this.parentElement.remove();
    }
    creating.append(cre, times);
    document.body.append(creating);
}

function practise() {
    let prac = document.getElementById("prac");
    prac.innerHTML = "{{.}}"
    console.log()
}