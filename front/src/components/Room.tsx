import React, { useState, useEffect } from 'react';

import { Link, useParams } from 'react-router-dom';
import { Button, Container, Group, List, Space, Paper, Text, TextInput } from '@mantine/core';

import { ChatService } from 'src/proto/proto/chat_connectweb.ts';
import { Message } from 'src/proto/proto/chat_pb.ts';
import { useClient } from 'src/client/client.ts';

const Room: React.FC = () => {
  const [messages, setMessages] = useState<Message[]>([]);
  const [text, setText] = useState<string>('');
  const [room, setRoom] = useState<string>('');
  const { id } = useParams<{ id: string }>();
  const client = useClient(ChatService);
  useEffect(() => { 
    const listMessages = async () => {
      const res = await client.listMessage({ roomId: id });
      setMessages(res.messages)
    }
    const getRoom = async () => {
      const res = await client.getRoom({ id: id });
      setRoom(res.room?.name || '');
    }
    listMessages();
    getRoom();
  }, []);

  const onClick = async () => {
    for await (const res of client.chat({
      message: {
        roomId: id,
        text: text,
      }
    })) {
      if (res.message) {
        setMessages([...messages, res.message]);
      }
    }
  }

  return (
    <Container size='xs'> 
      <Group position='right'>
        <Text>
          {room}
        </Text>
        <Button>
          <Link to='/' style={{ textDecoration: 'none' }}>
            <Text color='white'>ルーム一覧に戻る</Text>
          </Link>
        </Button>
      </Group>
      <List
        center
        spacing='xl'
        icon={<Space />}
      >
        {messages.map((message, i) => (
          <List.Item key={i}>
            <Text size='xl' align='center'>{message.text}</Text>
          </List.Item>
        ))}
      </List>
      <Paper shadow='md' m='xs' p='sm'>
        <TextInput mt='sm'
          label='メッセージ'
          withAsterisk
          placeholder='メッセージを入力してください'
          onChange={(e: React.ChangeEvent<HTMLInputElement>) => setText(e.target.value)}
        />
      </Paper>
      <Button mt='sm' onClick={onClick}>送信</Button>
    </Container>
  )
};

export default Room;
