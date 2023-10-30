// yo js :), Don't remove this comment.
const vocabs = document.querySelectorAll("[data-remove]");
const vocabForm = document.getElementById("vocab-form");
const ar = document.getElementById("arabic");
const eng = document.getElementById("english");
const harakat = document.getElementById("harakats");
const vocabFormErr = document.getElementById("vocab-form-err-div");


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
        vocabFormErr.innerHTML = `<span style="color:red">err: ${msg.err}</span>`;
    } else {
        if (res.redirected) {
            window.location.href = "/";
        }
    }
    reqesing = false;
})
