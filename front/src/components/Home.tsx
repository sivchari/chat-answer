import React, { useState, useEffect } from 'react';

import { Empty } from "@bufbuild/protobuf";
import { Button, Card, Container, Image, Group, List, Space, Text, Title } from '@mantine/core';
import { Link } from 'react-router-dom';

import { useClient } from 'src/client/client.ts';
import { Room } from 'src/proto/proto/chat_pb.ts';
import { ChatService } from 'src/proto/proto/chat_connectweb.ts';

const Home: React.FC = () => {
  const [rooms, setRooms] = useState<Room[]>([]);
  const client = useClient(ChatService);
  useEffect(() => {
    const listRooms = async () => {
      const res = await client.listRoom(Empty);
      setRooms(res.rooms);    
    };
    listRooms();
  }, []);
  return (
    <Container size='xs'>
      <Group position='right'>
        <Title order={2} align='center'>ルーム一覧</Title>
        <Button>
          <Link to='/room/create' style={{ textDecoration: 'none' }}>
            <Text color='white'>ルームを作成する</Text>
          </Link>
        </Button>
      </Group>
      <Space h='md' />
      <List
        center
        spacing='xl'
        icon={<Space />}
      >
        {rooms.map((room) => (
          <List.Item key={room.id}>
            <Card shadow='sm'>
              <Card.Section>
                <Image src='https://static.vecteezy.com/system/resources/previews/005/337/802/original/icon-symbol-chat-outline-illustration-free-vector.jpg' />
              </Card.Section>
              <Card.Section>
                <Text size='xl' align='center'>{room.name}</Text>
              </Card.Section>
              <Card.Section>
                <Button fullWidth>
                  <Link to={`/room/${room.id}`} style={{ textDecoration: 'none' }}>
                    <Text color='white'>入室する</Text>
                  </Link> 
                </Button>
              </Card.Section>
            </Card>
          </List.Item>
        ))}
      </List>
    </Container>
  )
};

export default Home;
