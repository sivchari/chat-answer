import React, { useState, useEffect, useRef } from 'react';

import { Link, useParams } from 'react-router-dom';
import { Button, Container, Group, List, ScrollArea, Space, Paper, Text, TextInput } from '@mantine/core';

import { ChatService } from 'src/proto/proto/chat_connectweb.ts';
import { Message } from 'src/proto/proto/chat_pb.ts';
import { useClient } from 'src/client/client.ts';

const Room: React.FC = () => {
  const [messages, setMessages] = useState<Message[]>([]);
  const [text, setText] = useState<string>('');
  const [room, setRoom] = useState<string>('');
  const { id } = useParams<{ id: string }>();
  const client = useClient(ChatService);
  const pass = useRef("");

  const joinRoom = async () => {
    const res = await client.getPass();
    pass.current = res.pass;

    console.log(pass.current, "join");
    for await (const res of client.joinRoom({ roomId: id, pass: pass.current })) {
      if (res.message) {
        setMessages(prev => [...prev, res.message]);
      }
    }
  }

  useEffect(() => { 
    const listMessages = async () => {
      const res = await client.listMessage({ roomId: id });
      setMessages(res.messages)
    }
    const getRoom = async () => {
      const res = await client.getRoom({ id: id });
      setRoom(res.room?.name || '');
    }
    const leaveRoom = async () => {
      console.log(pass.current, "leave");
      await client.leaveRoom({ roomId: id, pass: pass.current });
    }
    listMessages();
    getRoom();
    joinRoom();

    return () => {
      leaveRoom();
    }
  }, []);


  const handleClick = async () => {
    client.chat({
      message: {
        roomId: id,
        text: text,
      }
    }) 
    setText('');
  }

  const handleKeyPress = async (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      handleClick();
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
      <Space />
      <ScrollArea h='50vh' w='100%'>
        <List
          center
          spacing='md'
          icon={<Space />}
        >
          {messages.map((message, i) => (
            <List.Item key={i}>
              <Text size='xl' align='center'>{message.text}</Text>
            </List.Item>
          ))}
        </List>
      </ScrollArea> 
      <Paper shadow='md' m='xs' p='sm'>
        <Group>
          <TextInput
            value={text}
            onChange={(e) => setText(e.currentTarget.value)}
            onKeyPress={handleKeyPress}
            placeholder='メッセージを入力'
            pos='absolute'
            bottom={0}
            left={0}
            right={0}
          />
          <Button
            onClick={handleClick}
            pos='absolute'
            bottom={0}
            right={0}
          >
            送信
          </Button>
        </Group>
      </Paper>
    </Container>
  )
};

export default Room;
