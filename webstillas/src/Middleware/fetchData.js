import {BASE_URL} from "../Constants/apiURL"

/**
 * Function to fetch data from an api
 *
 * @param url we would like to fetch data from.
 * @returns {Promise<unknown>}
 */
export default function fetchModel(url) {
    return new Promise(function (resolve, reject) {
        const xhr = new XMLHttpRequest();

        xhr.addEventListener('readystatechange', () => {
            if (xhr.readyState !== 4) {
                return;
            }
            if (xhr.status !== 200) {
                reject({
                    status: xhr.status,
                    statusText: xhr.statusText,
                    text: xhr.responseText
                });
            } else {
                resolve({
                        statusCode: xhr.status,
                        text: xhr.responseText
                    }
                );
            }
        });
        xhr.open('GET',  BASE_URL + url);
        xhr.send();
    });
}



