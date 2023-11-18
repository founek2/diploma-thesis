import http from 'k6/http';
import { check, options } from 'k6';
import exec from 'k6/execution';

const API_URL = "http://localhost:8080/api/v1"

const fetchOptions = () => ({
    headers: {
        "User-Id": "user" + exec.vu.idInTest,
    }
});

function fetchItems() {
    const res = http.get(`${API_URL}/item`, fetchOptions());
    check(res, {
        'is status 200, fetchItems': (r) => r.status === 200,
    });

    return res;
}

/**
 * 
 * @param {string} itemId 
 * @returns 
 */
function addItemToCart(itemId) {
    const res = http.post(`${API_URL}/cart/items/${itemId}`, undefined, fetchOptions())
    check(res, {
        'is status 200, addItemToCart': (r) => r.status === 200,
    });

    return res;
}

function removeItemFromCart(itemId) {
    const res = http.del(`${API_URL}/cart/items/${itemId}`, undefined, fetchOptions())
    check(res, {
        'is status 200, removeItemFromCart': (r) => r.status === 200,
    });

    return res;
}

export default function () {
    const items = fetchItems().json();
    console.log(items)

    for (let i = 0; i < 100; i++) {
        const idx = Math.floor(Math.random() * items.length)
        addItemToCart(items[idx].id)
        removeItemFromCart(items[idx].id)
    }
}

export function handleSummary(data) {
    const d = new Date();

    return {
        [`results-latency/summary.json`]: JSON.stringify(data, null, 2), //the default data object
    };
}