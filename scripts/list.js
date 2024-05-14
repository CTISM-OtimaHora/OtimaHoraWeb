function adicionar() {
    const name = document.getElementById("new").value;
    document.getElementById("new").value = "";
    const obj = { id: undefined, nome: name };

    const params = new URLSearchParams(window.location.search)

    fetch(`http://localhost:3000/session/add/${params.get("tipo")}`, {
        method: "POST",
        credentials: "include",
        body: JSON.stringify(obj)
    })
    .then(response => response.text())
    .then(id => {
        document.getElementById("itens").appendChild(new_item(name, id));
    });
}

function new_item(name, id) {
    const params = new URLSearchParams(window.location.search)
    const tipo = params.get("tipo")

    const obj = document.createElement("div");
    obj.classList.add("item");

    const p_name = document.createElement("div");
    p_name.textContent = name;
    obj.appendChild(p_name);

    const btn = document.createElement("button");
    btn.onclick = () => {
        window.location.replace(`/dashboard.html?tipo=${tipo}&${tipo}=${name}&id=${id}`);
    };
    btn.classList.add('add')
    btn.textContent = "Editar";
    obj.appendChild(btn);

    return obj;
}

document.addEventListener("DOMContentLoaded", async () => {
    const tipo =  new URLSearchParams(window.location.search).get("tipo")
    document.getElementById("title").textContent = tipo
    document.getElementById("title2").textContent = `adicione ou modifique os ${tipo}`

    const res = await fetch(`http://localhost:3000/session/slice/${tipo}`, {
        credentials: "include"
    });
    const itens = await res.json();
    console.log(itens)

    for (const i of itens) {
        document.getElementById("itens").appendChild(new_item(
            i.Nome,
            i.Id
        ));
    }

});
