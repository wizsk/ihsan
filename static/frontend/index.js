// yo js :), Don't remove this comment.
// const okIcon = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-check-circle" viewBox="0 0 16 16"> <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"/> <path d="M10.97 4.97a.235.235 0 0 0-.02.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-1.071-1.05z"/> </svg>`;
const okIcon = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-check-circle-fill" viewBox="0 0 16 16"> <path d="M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0zm-3.97-3.03a.75.75 0 0 0-1.08.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-.01-1.05z"/> </svg>`; const vocabs = document.querySelectorAll("[data-remove]");
const vocabForm = document.getElementById("vocab-form");
const ar = document.getElementById("arabic");
const eng = document.getElementById("english");
const harakat = document.getElementById("harakats");
const vocabFormErr = document.getElementById("vocab-form-err-div");
const editVocabs = document.querySelectorAll("[data-edit]");


vocabs.forEach(e => {
    e.addEventListener("click", () => {
        const id = e.getAttribute("data-remove");
        fetch(`/api/remove?id=${id}`, {
            method: "POST",
        }).then((res) => {
            if (res.redirected) {
                // console.log(res.headers)
                window.location.href = "/"
            }
            console.log(res)
        }

        ).catch(err => {
            console.error(err);
        });
    })
});

editVocabs.forEach(e => {
    e.addEventListener("click", () => {
        const id = e.getAttribute("data-edit");
        const parent = document.getElementById(id);
        const icons = parent.querySelector(".icons");
        const ar = parent.querySelector(".arabic");
        const eng = parent.querySelector(".english");
        const arTxtPre = ar.innerText;
        const engTxtPre = eng.innerText;


        const keyDown = (e) => {
            if (e.keyCode == 13) {
                cleanAndReq()
            }
        };

        const cleanAndReq = async () => {
            let arTxt = ar.innerText;
            let engTxt = eng.innerText;
            arTxt = arTxt.trim();
            engTxt = engTxt.trim();
            console.log(ar, eng);
            console.log("ar", arTxtPre, arTxt, "eng", engTxtPre, engTxt);

            if (arTxt.includes("\n")) {
                console.error("ar text contins new line");
                cleanup();
                return;
            }
            if (engTxt.includes("\n")) {
                console.error("eng text contins new line");
                cleanup();
                return;
            }

            if (arTxtPre === arTxt && engTxt === engTxtPre) {
                console.log("nothing changed");
                cleanup();
                return;
            }

            console.log(arTxt, engTxt);

            const url = `/api/edit?id=${id}&arabic=${encodeURIComponent(arTxt)}&english=${encodeURIComponent(engTxt)}`;
            const res = await fetch(url, {
                method: "POST",
            });

            if (!res.ok) {
                console.error(res.error)
                cleanup(true)
                return;
            }

            cleanup();
            ar.innerText = arTxt;
            eng.innerText = engTxt;
        };

        // if default is ture then make text to pre.
        const cleanup = (defaut) => {
            if (defaut) {
                ar.innerText = arTxtPre;
                eng.innerText = engTxtPre;
            }

            // cleaning up
            ar.contentEditable = false;
            eng.contentEditable = false;
            icons.replaceChild(e, nd);
            ar.removeEventListener("keydown", keyDown);
            eng.removeEventListener("keydown", keyDown);
        };


        ar.addEventListener("keydown", keyDown)

        eng.addEventListener("keydown", keyDown);

        let nd = document.createElement("span");
        nd.innerHTML = okIcon;

        console.log(parent, e, nd)
        icons.replaceChild(nd, e);

        nd.addEventListener("click", cleanAndReq);


        ar.contentEditable = true;
        eng.contentEditable = true;

    })
});

let reqesing = false;
vocabForm.addEventListener("submit", async e => {
    e.preventDefault();
    if (reqesing) return;
    reqesing = true;

    let arabic = ar.value;
    arabic = arabic.trim();
    let english = eng.value;
    english = english.trim();
    // if (ar.value === "") {
    if (arabic === "") {
        vocabFormErr.innerHTML = `<span style="color:red">form value "arabic" is empty</span>`;
        return
        // } else if (eng.value === "") {
    } else if (english === "") {
        vocabFormErr.innerHTML = `<span style="color:red">form value "english" is empty</span>`;
        return
    }

    let respectHarakats = "false"
    if (harakat.checked) {
        respectHarakats = "true"
    }

    const url = `/api/add?arabic=${encodeURIComponent(arabic)}&english=${encodeURIComponent(english)}&respect_harakats=${respectHarakats}`
    const res = await fetch(url, {
        method: "POST",
    })

    if (!res.ok) {
        let msg = await res.json();
        let id = "";
        if (msg.data) {
            console.log(msg.data)
            id = `<br/> <a href="#${msg.data.id}">GoTo ${msg.data.arabic}</a>`
        }
        vocabFormErr.innerHTML = `<span style="color:red">err: ${msg.err}</span>${id}`;
    } else if (res.redirected) {
        window.location.href = "/";

    }
    reqesing = false;
})
