// Copyright (C) 2024 T-Force I/O
// This file is part of TF Unifiler
//
// TF Unifiler is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// TF Unifiler is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with TF Unifiler. If not, see <https://www.gnu.org/licenses/>.

package hasher

import "hash/crc32"

// HashCrc32 computes the CRC-32 checksum of the file at fPath using the IEEE polynomial.
func HashCrc32(fPath string) (*HashResult, error) {
	return hashFile(fPath, crc32.NewIEEE(), "crc32")
}

// HashCrc32c computes the CRC-32C checksum of the file at fPath using the Castagnoli polynomial.
func HashCrc32c(fPath string) (*HashResult, error) {
	return hashFile(fPath, crc32.New(crc32.MakeTable(crc32.Castagnoli)), "crc32c")
}
