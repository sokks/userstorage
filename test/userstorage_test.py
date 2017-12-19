import argparse
import json
import sys
import unittest

import requests

url = ""

class UserStorageTest(unittest.TestCase):

    headers = {'content-type': 'application/json'}

    # def setUp(self):

    @staticmethod
    def doRPC(method, args):
        payload = {
            "method": method,
            "params": [args], 
            "jsonrpc": "2.0",
            "id": 0,
        }
        response = requests.post(url, data=json.dumps(payload), headers=UserStorageTest.headers).json()
        return response
    
    def checkAddOK(self, login):
        response = UserStorageTest.doRPC("UserStorage.Add", {"login": login})
        self.assertEqual(0, response["id"])
        self.assertIsNone(response["error"])
        self.assertEqual("ok", response["result"]["Result"])

    def testNormalAdd(self):
        self.checkAddOK("testAddUserNo1")
    
    def testDoubleAdd(self):
        self.checkAddOK("testAddUserNo2")
        response = UserStorageTest.doRPC("UserStorage.Add", {"login": "testAddUserNo2"})
        self.assertEqual(0, response["id"])
        self.assertIsNotNone(response["error"])
        self.assertEqual("User already exists", response["error"])

    def testEmptyAdd(self):
        response = UserStorageTest.doRPC("UserStorage.Add", {"login": ""})
        self.assertEqual(0, response["id"])
        self.assertIsNotNone(response["error"])
        self.assertEqual("Incorrect login: empty", response["error"])

    def testLongAdd(self):
        longLogin = "testAddUserNo444444444444444444444444444444444444444"
        response = UserStorageTest.doRPC("UserStorage.Add", {"login": longLogin})
        self.assertEqual(0, response["id"])
        self.assertIsNotNone(response["error"])
        self.assertEqual("Incorrect login: too long", response["error"])

    def testNormalGet(self):
        self.checkAddOK("testGetUserNo1")
        response = UserStorageTest.doRPC("UserStorage.Get", {"login": "testGetUserNo1"})
        self.assertEqual(0, response["id"], 0)
        self.assertIsNone(response["error"])
        self.assertEqual("testGetUserNo1", response["result"]["Login"])
        self.assertTrue(response["result"]["UUID"])
        self.assertTrue(response["result"]["RegDate"])
        #self.assertRegex(response["result"]["UUID"], "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}")        

    def testNonexistingGet(self):
        response = UserStorageTest.doRPC("UserStorage.Get", {"login": "testGetUserNo2"})
        self.assertEqual(0, response["id"])
        self.assertIsNotNone(response["error"])
        self.assertEqual("User does not exist", response["error"])

    def testNormalRename(self):
        self.checkAddOK("testRenameUserNo1")
        response = UserStorageTest.doRPC("UserStorage.Rename", {'oldLogin': "testRenameUserNo1", "newLogin": "testRenameUserNo1New"})
        self.assertEqual(0, response["id"])
        self.assertIsNone(response["error"])
        self.assertEqual("ok", response["result"]["Result"])
        response = UserStorageTest.doRPC("UserStorage.Get", {"login": "testGetUserNo1"})
        self.assertEqual(0, response["id"])
        self.assertIsNone(response["error"])

    def testExistingRename(self):
        pass


urlPattern = "http://{host}:{port}/userstorage"
url = urlPattern.format(host='127.0.0.1', port='9081')

if __name__ == '__main__':
    unittest.main()
