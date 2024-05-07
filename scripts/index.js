document.addEventListener('DOMContentLoaded', async function() {
    await fetch("http://localhost:3000/add-session", {credentials:"include"})

    const res = await fetch("http://localhost:3000/session", {credentials:"include"})

    const sess = await res.json()
    
    let count = 0
    
    if (sess.Professores) {
        for (p of sess.Professores) {
            if (count == 3) {
                break
            }
            count += 1
            let child = document.createElement("li")
            child.textContent = p.Nome
            document.getElementById("professores").appendChild(child)
        }
    } else {
        document.getElementById("professores").appendChild(document.createTextNode("Nenhum professor criado"))
    }
    

    count = 0
    if (sess.Disciplinas) {
        for (d of sess.Disciplinas) {
            if (count == 3) {
                break
            }
            count += 1
            let child = document.createElement("li")
            child.textContent = d.Nome
            document.getElementById("disciplinas").appendChild(child)
        }
    } else {
        document.getElementById("disciplinas").appendChild(document.createTextNode("Nenhuma disciplina criada"))
    }

    count = 0
    if (sess.Cursos) {
        for (c of sess.Cursos) {
            if (count == 3) {
                break
            }
            count += 1
            let child = document.createElement("li")
            child.textContent = c.Nome
            document.getElementById("cursos").appendChild(child)
        }
    } else {
        document.getElementById("cursos").appendChild(document.createTextNode("Nenhum curso criado"))
    }

    count = 0
    if (sess.Contratos) {
        for (c of sess.Contratos) {
            if (count == 3) {
                break
            }
            count += 1
            let child = document.createElement("li")
            child.textContent = c.Nome
            document.getElementById("contratos").appendChild(child)
        }
    } else {
        document.getElementById("contratos").appendChild(document.createTextNode("Nenhum contrato criado"))
    }

});
