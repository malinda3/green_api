const defaultIdInstance = "1103157166";
const defaultApiTokenInstance = "cecd74e7d35849efa82c5a46f1fc543d618525b63cb84abda2";
const defaultChatId = "79687019003"; 
const defaultMessage = "test with message";
const defaultUrlFile = "https://sun9-37.userapi.com/impg/rWmGseJlnzKZjgbL9qNstDQyYpw7lo5S80Scgg/saRRNI0Crho.jpg?size=1170x1143&quality=95&sign=1b5e80e430c6338c90c9a4e06d2775b7&type=album"; 
const defaultCaption = "test with file"; 

function loadDefaultValues() {
    document.getElementById('idInstance').value = defaultIdInstance;
    document.getElementById('apiTokenInstance').value = defaultApiTokenInstance;
    document.getElementById('chatId').value = defaultChatId;
    document.getElementById('message').value = defaultMessage;
    document.getElementById('urlFile').value = defaultUrlFile;
    document.getElementById('caption').value = defaultCaption;
}

window.onload = loadDefaultValues;

const baseUrl = 'http://localhost:8881'; 

function normalizePhoneNumber(phone) {
    phone = phone.replace(/\D/g, '');

    if (phone.length === 11 && phone.startsWith('7')) {
        return phone + '@c.us';
    }

    if (phone.length === 12 && phone.startsWith('7')) {
        return phone.substring(1) + '@c.us';
    }

    throw new Error('Неверный формат номера телефона');
}

function getSettings() {
    const idInstance = document.getElementById('idInstance').value;
    const apiTokenInstance = document.getElementById('apiTokenInstance').value;
    console.log ("test:", idInstance, apiTokenInstance)
    fetch(`${baseUrl}/getSettings`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ idInstance, apiTokenInstance })
    })
    .then(response => {
        console.log(response);
        return response.text();
    })
    .then(data => {
        console.log(data); 
        try {
            const jsonData = JSON.parse(data);
            displayResponse(jsonData);
        } catch (error) {
            displayError('Ошибка при парсинге JSON:', error);
        }
    })
    .catch(error => displayError('Ошибка:', error));
}

function getStateInstance() {
    const idInstance = document.getElementById('idInstance').value;
    const apiTokenInstance = document.getElementById('apiTokenInstance').value;
    fetch(`${baseUrl}/getStateInstance`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ idInstance, apiTokenInstance })
    })
    .then(response => response.json())
    .then(data => {
        displayResponse(data);
    })
    .catch(error => displayError('Ошибка:', error));
}

function sendMessage() {
    let chatId = document.getElementById('chatId').value;
    const message = document.getElementById('message').value;
    const idInstance = document.getElementById('idInstance').value;
    const apiTokenInstance = document.getElementById('apiTokenInstance').value;

    try {
        chatId = normalizePhoneNumber(chatId);
    } catch (error) {
        displayError('Ошибка при нормализации номера:', error.message);
        return;
    }

    fetch(`${baseUrl}/send-message`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ chatId, message, idInstance, apiTokenInstance })
    })
    .then(response => response.json())
    .then(data => {
        displayResponse(data);
    })
    .catch(error => displayError('Ошибка:', error));
}

function sendFileByUrl() {
    let chatId = document.getElementById('fileChatId').value; 
    const urlFile = document.getElementById('urlFile').value; 
    const fileName = urlFile.split('/').pop().split('?')[0]; 
    const idInstance = document.getElementById('idInstance').value;
    const apiTokenInstance = document.getElementById('apiTokenInstance').value;

    try {
        chatId = normalizePhoneNumber(chatId);
    } catch (error) {
        displayError('Ошибка при нормализации номера:', error.message);
        return;
    }

    fetch(`${baseUrl}/send-file`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ chatId, urlFile, fileName, idInstance, apiTokenInstance })
    })
    .then(response => response.json())
    .then(data => {
        displayResponse(data);
    })
    .catch(error => displayError('Ошибка:', error));
}

function displayResponse(data) {
    const responseElement = document.getElementById('responseContent');
    const newResponse = JSON.stringify(data, null, 2);
    responseElement.textContent = newResponse;
}

function displayError(message, error) {
    const responseElement = document.getElementById('responseContent');
    responseElement.textContent = `${message} ${error}`;
}
