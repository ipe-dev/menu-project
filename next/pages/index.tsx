import { Flex, Heading, Input, Button, useColorMode } from '@chakra-ui/react'
import withSession

const indexPage = () => {
  return (
    <Flex height="100vh" alignItems="center" justifyContent="center">
      <Flex direction="column" background="gray.100" p={12} rounded={6} >
        <Heading mb={6}>Top</Heading>
     </Flex>
    </Flex>

  )
}
export const getServerSideProps = withSession(async function ({req, res}) {
  const user = req.session

  if (!user) {
    return {
      redirect: {
        destination: "/login",
        permanent: false
      }
    }
  }
  return {
    props: user
  }
})
  

export default indexPage
