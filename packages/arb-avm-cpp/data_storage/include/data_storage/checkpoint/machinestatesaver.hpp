/*
 * Copyright 2019, Offchain Labs, Inc.
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

#ifndef machinestatesaver_hpp
#define machinestatesaver_hpp

#include <avm_values/tuple.hpp>
#include <avm_values/value.hpp>

struct MachineStateKeys;

class Transaction;
struct SaveResults;

namespace rocksdb {
class Status;
}

class MachineStateSaver {
   private:
    std::unique_ptr<Transaction> transaction;

   public:
    MachineStateSaver(std::unique_ptr<Transaction> transaction_);
    SaveResults saveTuple(const Tuple& val);
    SaveResults saveValue(const value& val);
    SaveResults saveMachineState(
        const MachineStateKeys& state_data,
        const std::vector<unsigned char>& checkpoint_name);
    rocksdb::Status commitTransaction();
    rocksdb::Status rollBackTransaction();
};

#endif /* machinestatesaver_hpp */
