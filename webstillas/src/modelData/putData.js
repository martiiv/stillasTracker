import {BASE_URL} from "./constantsFile";

export default function putModel(url, body) {
    return new Promise(function (resolve, reject) {
        const xhr = new XMLHttpRequest();
        xhr.open("PUT", BASE_URL + url);
        xhr.setRequestHeader("Content-Type", "application/json");
        /*
            load event is also ok to use here,
            but readystatechange was giving me more descriptive errors
        */
        xhr.addEventListener('readystatechange', () => {
            if (xhr.readyState !== 4) {
                return;
            }
            if (xhr.status !== 200) {
                reject((({
                    statusCode: xhr.status,
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
        xhr.send(body);
    });
}
