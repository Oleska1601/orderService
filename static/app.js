function fetchOrder() {
            const orderId = document.getElementById('orderId').value.trim();
            const messageElement = document.getElementById('message');
            const orderInfoElement = document.getElementById('orderInfo');
            
            // сбрасываются предыдущ сообщения и скрываются результаты
            messageElement.textContent = '';
            messageElement.className = 'error-message';
            orderInfoElement.style.display = 'none';
            
            if (!orderId) {
                messageElement.textContent = 'No order UID provided';
                return;
            }
            
            fetch(`http://localhost:8081/order/${orderId}`)
                .then(response => {
                    if (response.status === 400 || response.status === 404 || response.status === 500) {
                        return response.text().then(text => { throw new Error(text); });
                    }
                    if (response.ok) {
                        return response.json();
                    }
                    throw new Error('Unknown error occurred');
                })
                .then(data => {
                    messageElement.textContent = 'Order loaded successfully';
                    messageElement.className = 'success-message';
                    displayOrder(data);
                })
                .catch(error => {
                    messageElement.textContent = error.message;
                    messageElement.className = 'error-message';
                });
        }
        
function displayOrder(order) {
    document.getElementById('orderInfo').style.display = 'block';
    const orderData = {...order};
    delete orderData.delivery;
    delete orderData.payment;
    delete orderData.items;
            
    document.getElementById('orderData').textContent = JSON.stringify(orderData, null, 2); // (value, replacer, space)
    document.getElementById('deliveryData').textContent = JSON.stringify(order.delivery, null, 2);
    document.getElementById('paymentData').textContent = JSON.stringify(order.payment, null, 2);
    document.getElementById('itemsData').textContent = JSON.stringify(order.items, null, 2);
}