function cadastre() {
    let idbutton = document.getElementsByClassName("upload");
    idbutton[0].remove();
    idbutton[0].remove();
    let parag = document.getElementById("error");
    parag.remove();

    let createform = document.createElement("form");
    createform.setAttribute("action", "");
    createform.setAttribute("enctype", "multipart/form-data");
    createform.setAttribute("method", "POST");


    let input = document.createElement("input");
    input.setAttribute("type", "file");
    input.setAttribute("id", "form-button");
    input.setAttribute("name", "upload");
    input.hidden = true;
    let labeler = document.createElement("label");
    labeler.setAttribute("for", "form-button");
    let upload = document.createElement("i");
    upload.classList.add("fas", "fa-file-upload");
    labeler.append(upload);
    let number = document.createElement("input");
    number.setAttribute("name", "sheetnumber");
    number.setAttribute("type", "number");
    number.value = 0;
    let submit = document.createElement("input");
    submit.setAttribute("type", "submit");
    submit.setAttribute("value", "Upload");



    createform.append(input, labeler, number, submit);

    document.body.append(createform);

}