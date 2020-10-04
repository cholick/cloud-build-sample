import logging
import sys
import unittest

import requests


class TestE2E(unittest.TestCase):
    def setUp(self):
        self.host = "staging.example.com"
        self.log = logging.getLogger(__name__, )
        logging.basicConfig(stream=sys.stderr, level=logging.DEBUG)

    def test_works(self):
        url = "http://{}/foo".format(self.host)
        resp = requests.get(url)

        self.assertTrue(resp.ok, "Response status not ok")
        self.assertIn("Success", resp.text, "Body did not contain string 'Success'")
