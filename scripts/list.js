function adicionar() {
    const name = document.getElementById("new").value;
    if (!name || name === "") {
        return
    }
    document.getElementById("new").value = "";
    const obj = { id: undefined, nome: name };

    const params = new URLSearchParams(window.location.search)

    fetch(`http://localhost:3000/${params.get("tipo")}/add`, {
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
    p_name.classList.add("nome")

    obj.appendChild(p_name);

    const btn = document.createElement("button");
    btn.onclick = () => {
        window.location.replace(`/OtimaHoraWeb/dashboard.html?tipo=${tipo}&${tipo}=${name}&id=${id}`);
    };
    btn.classList.add('add')
    btn.textContent = "Editar";
    const rem_btn = document.createElement("button");
    rem_btn.onclick = () => {
        fetch(`http://localhost:3000/${params.get("tipo")}/delete/${id}`, {credentials: "include", method:"DELETE"}).then(obj.parentNode.removeChild(obj))
    };
    rem_btn.classList.add('rem')
    rem_btn.textContent = "Remover";

    obj.appendChild(btn)
    obj.appendChild(rem_btn);

    return obj;
}

document.addEventListener("DOMContentLoaded", async () => {
    const tipo =  new URLSearchParams(window.location.search).get("tipo")
    document.getElementById("title2").textContent = `Alterar ${tipo}:`

    const res = await fetch(`http://localhost:3000/${tipo}/slice`, {
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

    document.addEventListener("keypress", (key) => {
            if (key.key === "Enter") {
            document.getElementById("adicionar").click()
        }
    })

});

