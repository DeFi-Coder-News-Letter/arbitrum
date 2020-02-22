/*
 * Copyright 2019-2020, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

const SigUtilsTester = artifacts.require("SigUtilsTester");

contract("SigUtils", accounts => {
  it("it should get recover address from sig", async () => {
    let sig_utils_tester = await SigUtilsTester.new();
    let raw_message = web3.utils.soliditySha3("0x1234");
    let eth_message = web3.eth.accounts.hashMessage(raw_message);
    let sig = await web3.eth.sign(raw_message, accounts[0]);
    let addr = await sig_utils_tester.recoverAddress(raw_message, sig);
    assert.equal(addr, accounts[0]);
  });

  it("it should get recover address from offset sig", async () => {
    let sig_utils_tester = await SigUtilsTester.new();
    let raw_message1 = web3.utils.soliditySha3("0x1234");
    let sig1 = await web3.eth.sign(raw_message1, accounts[0]);
    let raw_message2 = web3.utils.soliditySha3("0x4321");
    let sig2 = await web3.eth.sign(raw_message2, accounts[0]);
    let combo_sig = sig1 + sig2.slice(2);
    let addr1 = await sig_utils_tester.recoverAddress(
      raw_message1,
      combo_sig,
      0
    );
    let addr2 = await sig_utils_tester.recoverAddress(
      raw_message2,
      combo_sig,
      1
    );
    assert.equal(addr1, accounts[0]);
    assert.equal(addr2, accounts[0]);
  });
});
