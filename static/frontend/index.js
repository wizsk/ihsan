// yo js :), Don't remove this comment.
const vocabs = document.querySelectorAll("[data-remove]");
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
        const ar = parent.querySelector(".arabic");
        const eng = parent.querySelector(".english");
        const arTxtPre = ar.innerText;
        const engTxtPre = ar.innerText;

        const keyDown = (e) => {
            if (e.keyCode == 13) {
                cleanAndReq()
            }
        };

        const cleanAndReq = () => {
            // cleaning up
            ar.contentEditable = false;
            eng.contentEditable = false;
            parent.replaceChild(e, nd);
            ar.removeEventListener("keydown", keyDown);
            eng.removeEventListener("keydown", keyDown);

            let arTxt = ar.innerText;
            let engTxt = eng.innerText;
            arTxt = arTxt.trim();
            engTxt = engTxt.trim();

            if (arTxt.includes("\n")) {
                console.error("ar text contins new line");
                return;
            }
            if (engTxt.includes("\n")) {
                console.error("eng text contins new line");
                return;
            }

            console.log(arTxt, engTxt);
        };


        ar.addEventListener("keydown", keyDown)

        eng.addEventListener("keydown", keyDown);

        let nd = document.createElement("td");
        nd.innerText = "submit ->"

        parent.replaceChild(nd, e);

        nd.addEventListener("click", cleanAndReq);


        ar.contentEditable = true;
        eng.contentEditable = true;


        // fetch(`/api/remove?id=${id}`, {
        //     method: "POST",
        // }).then((res) => {
        //     if (res.redirected) {
        //         // console.log(res.headers)
        //         window.location.href = "/"
        //     }
        //     console.log(res)
        // }

        // ).catch(err => {
        //     console.error(err);
        // });
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
    } else {
        if (res.redirected) {
            window.location.href = "/";
        }
    }
    reqesing = false;
})
