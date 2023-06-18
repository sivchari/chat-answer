import React, { useState, useEffect } from 'react';

import { Link, useParams } from 'react-router-dom';
import { Button, Container, Group, List, Space, Text } from '@mantine/core';

import { ChatService } from 'src/proto/proto/chat_connectweb.ts';
import { Message } from 'src/proto/proto/chat_pb.ts';
import { useClient } from 'src/client/client.ts';

const Room: React.FC = () => {
  const [messages, setMessages] = useState<Message[]>([]);
  const [room, setRoom] = useState<string>('');
  const { id } = useParams<{ id: string }>();
  useEffect(() => { 
    const client = useClient(ChatService);
    const listMessages = async () => {
      const res = await client.listMessage({ roomId: id });
      setMessages(res.messages)
    }
    const getRoom = async () => {
      const res = await client.getRoom({ id: id });
      setRoom(res.name);
    }
    listMessages();
    getRoom();
  }, []);
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
    </Container>
  )
};

export default Room;
