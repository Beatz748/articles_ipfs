use mydb

// Создание коллекции и вставка данных
db.articles.insertMany([
    {
        content: 'Это содержание статьи 1.',
        _id: '1'
    },
    {
        content: 'Это содержание статьи 2.',
        _id: '2'
    }
])

print('Инициализация MongoDB завершена.')

