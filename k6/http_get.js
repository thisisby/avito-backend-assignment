import http from 'k6/http';

export const options = {
    vus: 1,
    duration: '10s',
}

export default function () {
    const url = 'http://localhost:8080/api/v1/token/generate';
    const payload = JSON.stringify({
        type: 1,
        length: 18,
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.post(url, payload, params);
}