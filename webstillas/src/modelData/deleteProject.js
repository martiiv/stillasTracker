import {BASE_URL} from "./constantsFile"

export default function deleteModel(url, body) {
    return new Promise(function (resolve, reject) {
        const xhr = new XMLHttpRequest();
        /*
            load event is also ok to use here,
            but readystatechange was giving me more descriptive errors
        */
        xhr.addEventListener('readystatechange', () => {
            if (xhr.readyState !== 4) {
                return;
            }
            if (xhr.status !== 200) {
                reject(new Error(JSON.stringify({
                    status: xhr.status,
                    statusText: xhr.statusText,
                    text: xhr.responseText
                })));
            } else {
                resolve((xhr.responseText));
            }
        });
        xhr.open('DELETE',  BASE_URL + url);
        xhr.send(JSON.stringify(body));
    });
}
