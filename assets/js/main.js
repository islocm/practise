function OpenWindowLogin() {
    let modal = document.createElement("div");
    modal.className = "user-login";


    let inputLogin = document.createElement("input");
    inputLogin.type = "text";
    inputLogin.placeholder = "Login";

    let inputPass = document.createElement("input");
    inputPass.type = "password";
    inputPass.placeholder = "Password";

    let button = document.createElement("button");
    button.onclick = Loginquery.bind(button, inputLogin, inputPass);
    button.textContent = "Login";

    let closeWindow = document.createElement("i");
    closeWindow.classList.add("fas", "fa-times");
    closeWindow.onclick = function() {
        this.parentElement.parentElement.remove();
    }

    let background = document.createElement("div");
    background.className = "model";

    modal.append(inputLogin, inputPass, button, closeWindow);
    background.append(modal);
    document.body.append(background);



}


function Loginquery(login, pass) {
    let isCorrect = true

    if (!login || login.value.length < 4) {
        login.classList.add("incorrect");
        login.oninput = Clearincorrect;
        isCorrect = false
    }

    if (!pass || pass.value.length < 3) {
        pass.classList.add("incorrect");
        pass.oninput = Clearincorrect;
        isCorrect = false
    }

    if (!isCorrect) {
        return;
    }

    let xhr = new XMLHttpRequest();

    xhr.open("POST", "/login")

    xhr.onload = function(event) {
        try {
            let data = JSON.parse(event.target.response)

            if (data && "Error" in data && data.Error == null) {
                if ("Name" in data) {
                    window.location.reload()
                }
            } else {
                console.log(data.Error)
                alert(data.Error)
            }

        } catch (error) {
            console.log(error)
        }
    }
    let data = JSON.stringify({
        Tel: login.value,
        Password: pass.value,
    })
    console.log(data);
    xhr.send(data)
}

function Clearincorrect() {
    this.classList.remove("incorrect");
    this.oninput = undefined;
}