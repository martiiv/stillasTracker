import {BASE_URL} from "./constantsFile"
function fetchModel(url) {
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
                resolve(JSON.parse(xhr.responseText));
            }
        });
        xhr.open('GET',  BASE_URL + url);
        xhr.send();
    });
}

export default fetchModel
