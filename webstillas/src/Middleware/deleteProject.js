import {BASE_URL} from "../Constants/apiURL"

/**
 * Function that will send a delete request to an api.
 *
 * @param url to the request we would like to send
 * @param body request body
 * @returns {Promise<unknown>}
 */
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
                resolve({
                        statusCode: xhr.status,
                        text: xhr.responseText
                    }
                );
            }
        });
        xhr.open('DELETE',  BASE_URL + url);
        xhr.send(JSON.stringify(body));
    });
}
