import socket;

conn = socket.socket();

conn.connect(("localhost",8001));

data = "";
while data != "q":
  data = input("Enter a Message> ")
  conn.send(data.encode(encoding='utf_8',errors='strict'))
