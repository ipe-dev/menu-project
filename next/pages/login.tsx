import { Flex, Heading, Input, Button } from '@chakra-ui/react'

const Login = () => {
  const login = () => {
    
  }
  return (
    <Flex height="100vh" alignItems="center" justifyContent="center">
      <Flex direction="column" background="gray.100" p={12} rounded={6} >
        <Heading mb={6}>Log in</Heading>
        <Input placeholder="hoge@chakra-ui.com" variant="filled" mb={3} type="email" />
        <Input placeholder="*******" variant="filled" mb={6} type="password" />
        <Button mb={6} colorScheme="teal">Log in</Button>
      </Flex>
    </Flex>
  )
}

export default Login