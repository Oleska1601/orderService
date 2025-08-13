document.getElementById('searchButton').addEventListener('click', getOrderInfo);

async function getOrderInfo() {
    const orderId = document.getElementById('order_id').value.trim();
    const resultDiv = document.getElementById('result');
    const errorDiv = document.getElementById('error');
    
    resultDiv.style.display = 'none';
    errorDiv.textContent = '';
    
    if (!orderId) {
        errorDiv.textContent = 'Пожалуйста, введите ID заказа';
        return;
    }
    
    try {
        const response = await fetch(`http://localhost:8081/order/${orderId}`);
        
        if (!response.ok) {
            throw new Error('Заказ не найден');
        }
        
        const order = await response.json();
        displayOrderInfo(order);
        resultDiv.style.display = 'block';
    } catch (error) {
        errorDiv.textContent = error.message;
    }
}

function displayOrderInfo(order) {
    fetch('order-info.html')
        .then(response => response.text())
        .then(template => {
            const resultDiv = document.getElementById('result');
            resultDiv.innerHTML = template;
            
            // Заполняем шаблон данными
            document.getElementById('orderBasicInfo').innerHTML = `
                <p><strong>ID заказа:</strong> ${order.order_uid}</p>
                <p><strong>Трек-номер:</strong> ${order.track_number}</p>
                <p><strong>Дата создания:</strong> ${new Date(order.date_created).toLocaleString()}</p>
                <p><strong>Служба доставки:</strong> ${order.delivery_service}</p>
                <p><strong>Клиент:</strong> ${order.customer_id}</p>
            `;
            
            // Остальные части заполняются аналогично...
            // Для полноты нужно добавить заполнение всех секций
        });
}

function getStatusText(status) {
    const statuses = {
        202: 'Принят',
        203: 'В обработке',
        204: 'Отправлен',
        205: 'Доставлен',
        206: 'Отменен'
    };
    return statuses[status] || `Неизвестный статус (${status})`;
}