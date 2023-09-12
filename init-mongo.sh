#!/bin/bash

# Ожидайте доступности MongoDB
until mongosh -u root -p qwerty --eval "print(\"waited for connection\")" 2>/dev/null; do
  echo "MongoDB is not available yet. Retrying in 1 second..."
  sleep 1
done

# Создайте базу данных mydb
mongosh -u root -p qwerty --eval "db.getSiblingDB('papers')"

# Создайте коллекцию articles и вставьте данные
mongosh -u root -p qwerty --authenticationDatabase admin papers <<EOF
db.articles.insertMany([
    {
        content: 'Blockchains are distributed (i.e., without a single repository) and decentralized digital ledgers that are tamper-evident and resistant. At their most basic level, they allow users to record transactions in a shared ledger within that group. The result is that no transaction can be modified once it has been published under standard blockchain network functioning. The blockchain concept was integrated with numerous other technologies and computer concepts in 2008 to create modern cryptocurrencies: electronic cash that is protected by cryptographic processes rather than a central repository or authority.',
        "_id": '1'
    },
    {
        content: 'Это содержание статьи 2.',
        "_id": '2'
    }
]);
EOF

echo "Инициализация MongoDB завершена."

