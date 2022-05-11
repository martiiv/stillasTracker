import {BASE_URL} from "./constantsFile";

export default function postModel(url, body) {
    return new Promise(function (resolve, reject) {
        const xhr = new XMLHttpRequest();
        xhr.open("POST", BASE_URL + url);
        xhr.setRequestHeader("Content-Type", "application/json");
        /*
            load event is also ok to use here,
            but readystatechange was giving me more descriptive errors
        */
        xhr.addEventListener('readystatechange', () => {
            if (xhr.readyState !== 4) {
                return;
            }
            if (xhr.status !== 201) {
                reject(new Error(JSON.stringify({
                    status: xhr.status,
                    statusText: xhr.statusText,
                    text: xhr.responseText
                })));
            } else {
                resolve({
                    statusCode: xhr.status,
                    text: xhr.responseText
                    }
                );
            }
        });

        xhr.send(JSON.stringify(body));
    });
}
