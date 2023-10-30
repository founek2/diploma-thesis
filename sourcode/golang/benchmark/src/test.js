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

function createOrder() {
    const res = http.post(`${API_URL}/order/create`, undefined, fetchOptions())
    check(res, {
        'is status 200, createOrder': (r) => r.status === 200,
    });

    return res;
}

function cancelOrder(orderId) {
    const res = http.post(`${API_URL}/order/${orderId}/cancel`, undefined, fetchOptions())
    check(res, {
        'is status 200, cancelOrder': (r) => r.status === 200,
    });

    return res;
}

function fetchInvoice(invoiceId) {
    const res = http.get(`${API_URL}/invoice/${invoiceId}`, fetchOptions())
    check(res, {
        'is status 200, fetchInvoice': (r) => r.status === 200,
    });

    return res;
}

function payForInvoice(invoiceId) {
    const res = http.post(`${API_URL}/payment/invoice/${invoiceId}`, undefined, fetchOptions())
    check(res, {
        'is status 200, pay': (r) => r.status === 200,
    });

    return res;
}

export default function () {
    const items = fetchItems().json();

    for (let i = 0; i < 10; i++) {
        const idx = Math.floor(Math.random() * items.length)
        addItemToCart(items[idx].id)
    }

    const order = createOrder().json()
    const invoice = fetchInvoice(order.invoiceId).json()
    // cancelOrder(order.orderId)
    payForInvoice(invoice.id)
}