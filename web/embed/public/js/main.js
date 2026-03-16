new class {
    constructor(data) {
        this.data = data;
        this.data.input.addEventListener("input", () => {
            if (this.data["input"].value.length > 0) {
                this.hideDescriptionTextContainer();
                this.showFooterTextContainer();
            } else {
                this.showDescriptionTextContainer();
                this.hideFooterTextContainer();
            };
            this.hideErrorText();
            this.hideSection_2Container();
            this.hideSection_3Container();
            this.hideSection_4Container();
        });
        this.data.postButton.addEventListener("click", () => {
            if (!this.data.input.value.trim().length === 0 || !this.data.input.value.includes(".")) {
                this.showErrorText("Некорректная ссылка");
                this.hideFooterTextContainer();
            } else {
                this.post(this.data.input.value);
            };
        });
        this.data.copyButton_1.addEventListener("click", () => this.copyLink(this.data.linkResultURL.href));
        this.data.copyButton_2.addEventListener("click", () => this.copyLink(this.data.linkPageURL.href));
        this.hideErrorText();
        this.hideFooterTextContainer();
        this.hideSection_2Container();
        this.hideSection_3Container();
        this.hideSection_4Container();
        if (document.location.search.length > 0 && document.location.search[0] === "?") {
            var args = document.location.search.split("=")
            if (args.length === 2 && args[0].length === 2 && args[0][1] == "s" && args[1].length > 0) {
                this.get(args[1]);
            };
        };
    };
    showErrorText(text) {
        this.data.errorTextContainer.classList.add("default-show");
        this.data.errorTextContainer.classList.remove("default-hide");
        this.data.errorText.textContent = text;
    };
    hideErrorText() {
        this.data.errorTextContainer.classList.add("default-hide");
        this.data.errorTextContainer.classList.remove("default-show");
        this.data.errorText.textContent = "";
    };
    showDescriptionTextContainer() {
        this.data.descriptionTextContainer.classList.add("default-show");
        this.data.descriptionTextContainer.classList.remove("default-hide");
    };
    hideDescriptionTextContainer() {
        this.data.descriptionTextContainer.classList.add("default-hide");
        this.data.descriptionTextContainer.classList.remove("default-show");
    };
    showFooterTextContainer() {
        this.data.footerTextContainer.classList.add("default-show");
        this.data.footerTextContainer.classList.remove("default-hide");
    };
    hideFooterTextContainer() {
        this.data.footerTextContainer.classList.add("default-hide");
        this.data.footerTextContainer.classList.remove("default-show");
    };
    showSection_2Container() {
        this.data.section_2Container.classList.add("default-show");
        this.data.section_2Container.classList.remove("default-hide");
    };
    hideSection_2Container() {
        this.data.section_2Container.classList.add("default-hide");
        this.data.section_2Container.classList.remove("default-show");
    };
    showSection_3Container() {
        this.data.section_3Container.classList.add("default-show");
        this.data.section_3Container.classList.remove("default-hide");
    };
    hideSection_3Container() {
        this.data.section_3Container.classList.add("default-hide");
        this.data.section_3Container.classList.remove("default-show");
    };
    showSection_4Container() {
        this.data.section_4Container.classList.add("default-show");
        this.data.section_4Container.classList.remove("default-hide");
    };
    hideSection_4Container() {
        this.data.section_4Container.classList.add("default-hide");
        this.data.section_4Container.classList.remove("default-show");
    };
    copyLink(url) {
        try {
            navigator.clipboard.writeText(url.replace("xn--q1a", "с")).then(() => {
                alert("Скопировано!");
            }).catch(error => {
                alert(`Ошибка копирования: ${error}`)
            });
        } catch (error) {
            alert(`Ошибка копирования: ${error}`)
        };
    };
    setOriginalLink(url) {
        this.data.input.value = url;
    };
    setLinesContainer(clicks, created_at, last_accessed_at) {
        var data = [
            ["Нажатий:", clicks],
            ["Создана:", this.liveTimeAgo(created_at)],
            ["Последний клик:", last_accessed_at ? this.liveTimeAgo(last_accessed_at) : "никогда"],
        ];
        linesContainer.innerHTML = "";
        for (var i = 0; i < data.length; i++) {
            const div = document.createElement("div");
            const p = document.createElement("p");
            div.className = "shortener__result-line-container";
            p.textContent = data[i][0] + " ";
            if (data[i][1] instanceof HTMLElement) {
                p.appendChild(data[i][1]);
            } else {
                const span = document.createElement("span");
                span.textContent = data[i][1];
                p.appendChild(span);
            }
            div.appendChild(p);
            linesContainer.appendChild(div);
            if (i < data.length - 1) {
                linesContainer.appendChild(
                    document.createElement("hr"),
                );
            };
        };
    };
    setResultLink(short_url, short_code, domain) {
        this.data.linkResultText.textContent = `${domain}/${short_code}`;
        this.data.linkResultURL.href = short_url;
    };
    setPageLink(short_code, domain) {
        this.data.linkPageText.textContent = `${domain}/?s=${short_code}`;
        this.data.linkPageURL.href = `https://${domain}/?s=${short_code}`;;
    };
    liveTimeAgo(dateString) {
        const span = document.createElement("span");
        const start = new Date(dateString);
        const units = [
            ["год", 31536000],
            ["день", 86400],
            ["час", 3600],
            ["минута", 60],
            ["секунда", 1],
        ];
        function plural(n, word) {
            const forms = {
                "секунда": ["секунду", "секунды", "секунд"],
                "минута": ["минуту", "минуты", "минут"],
                "час": ["час", "часа", "часов"],
                "день": ["день", "дня", "дней"],
                "год": ["год", "года", "лет"],
            };
            const f = forms[word];
            if (n % 10 === 1 && n % 100 !== 11) return f[0];
            if (n % 10 >= 2 && n % 10 <= 4 && (n % 100 < 10 || n % 100 >= 20)) return f[1];
            return f[2];
        };
        function update() {
            let diff = Math.floor((new Date() - start) / 1000);
            if (diff < 0) {
                span.textContent = "неизвестно";
                return;
            };
            if (diff < 1) {
                span.textContent = "только что";
                return;
            };
            for (const [name, sec] of units) {
                const value = Math.floor(diff / sec);
                if (value >= 1) {
                    span.textContent = `${value} ${plural(value, name)} назад`;
                    return;
                };
            };
        };
        update();
        setInterval(update, 1000);
        return span;
    };
    async post(url) {
        try {
            const response = await fetch("/api/links", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    "url": url,
                }),
            });
            if (response.ok) {
                document.location.href = `/?s=${(await response.json())["short_code"]}`;
            } else {
                this.showErrorText(`Ошибка с сервера: ${await response.text()}`);
                this.hideFooterTextContainer();
            }
        } catch {
            this.showErrorText("Сервер не ответил");
            this.hideFooterTextContainer();
        };
    };
    async get(short_code) {
        try {
            const response = await fetch(`/api/links/${short_code}`, {
                method: "GET",
            });
            if (response.ok) {
                this.hideDescriptionTextContainer();
                this.showSection_2Container();
                this.showSection_3Container();
                this.showSection_4Container();
                ((response_data) => {
                    this.setOriginalLink(
                        response_data["original_url"],
                    );
                    this.setLinesContainer(
                        response_data["clicks"],
                        response_data["created_at"],
                        response_data["last_accessed_at"],
                    );
                    this.setResultLink(
                        response_data["short_url"],
                        response_data["short_code"],
                        response_data["domain"],
                    );
                    this.setPageLink(
                        response_data["short_code"],
                        response_data["domain"],
                    );
                })(await response.json());
            } else {
                document.location.href = "/";
            }
        } catch (error) {
            document.location.href = "/";
        };
    };
}({
    input: document.getElementById("inp"),
    postButton: document.getElementById("postBtn"),
    copyButton_1: document.getElementById("copyBtn_1"),
    copyButton_2: document.getElementById("copyBtn_2"),
    errorText: document.getElementById("errorText"),
    errorTextContainer: document.getElementById("errorTextContainer"),
    descriptionTextContainer: document.getElementById("descriptionTextContainer"),
    footerTextContainer: document.getElementById("footerTextContainer"),
    section_2Container: document.getElementById("section_2Container"),
    linkResultURL: document.getElementById("linkResultURL"),
    linkResultText: document.getElementById("linkResultText"),
    section_3Container: document.getElementById("section_3Container"),
    linesContainer: document.getElementById("linesContainer"),
    section_4Container: document.getElementById("section_4Container"),
    linkPageURL: document.getElementById("linkPageURL"),
    linkPageText: document.getElementById("linkPageText"),
});
