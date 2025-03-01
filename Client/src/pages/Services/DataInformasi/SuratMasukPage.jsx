import React, { useState, useEffect } from "react";
import App from "../../../components/Layouts/App";
import Swal from "sweetalert2";
import { ReusableForm } from "../../../components/Fragments/Services/ReusableForm";
import { ReusableTable } from "../../../components/Fragments/Services/ReusableTable";
import { jwtDecode } from "jwt-decode";
import { Modal } from "flowbite-react";
import {
  getSuratMasuks,
  addSuratMasuk,
  deleteSuratMasuk,
  updateSuratMasuk,
} from "../../../../API/DataInformasi/SuratMasuk.service";

export function SuratMasukPage() {
  const [MainData, setMainData] = useState([]);
  const [formModalOpen, setFormModalOpen] = useState(false);
  const [formData, setFormData] = useState({});
  const [formConfig, setFormConfig] = useState({
    fields: [
      { name: "no_surat", label: "No Surat", type: "text", required: true },
      { name: "title", label: "Title Of Letter", type: "text", required: true },
      {
        name: "related_div",
        label: "Related Divisi",
        type: "text",
        required: true,
      },
      { name: "destiny_div", label: "Destiny", type: "text", required: true },
      { name: "tanggal", label: "Date Issue", type: "date", required: true },
    ],
    services: "Surat Masuk",
  });
  const [selectedIds, setSelectedIds] = useState([]);
  const token = localStorage.getItem("token");
  let userRole = "";
  if (token) {
    const decoded = jwtDecode(token);
    userRole = decoded.role;
  }

  // UseEffect untuk mengambil data saat komponen dimount dan di balikan urutan
  useEffect(() => {
    getSuratMasuks((data) => {
      // ambil dari API
      setMainData(data.reverse());
    });
  }, []);

  // Function untuk fetch data dan update state
  const handleAdd = () => {
    setFormModalOpen(true);
    setFormConfig((prevConfig) => ({
      ...prevConfig,
      action: "add",
      onSubmit: (data) => AddSubmit(data),
    }));
  };

  // Function untuk handle tutup form modal
  const onCloseFormModal = () => {
    setFormModalOpen(false);
    setFormData({});
  };

  // Function untuk fetch data dan update state
  const handleEdit = (MainData) => {
    setFormModalOpen(true);
    setFormConfig((prevConfig) => ({
      ...prevConfig,
      action: "edit",
      onSubmit: (data) => EditSubmit(data),
    }));
    setFormData({ ...MainData });
  };

  // tambah data
  const AddSubmit = async (data) => {
    try {
      await addSuratMasuk(data); // tambah data ke API
      Swal.fire({
        icon: "success",
        title: "Berhasil!",
        text: "Data berhasil ditambahkan",
        showConfirmButton: false,
        timer: 1500,
      }).then(() => {
        window.location.reload();
        setMainData([...MainData, data]);
      });
    } catch (error) {
      Swal.fire({
        icon: "error",
        title: "Gagal!",
        text: "Error saat menyimpan data",
        showConfirmButton: false,
        timer: 1500,
      });
    } finally {
      onCloseFormModal();
    }
  };

  // ubah data
  const EditSubmit = async (data) => {
    try {
      await updateSuratMasuk(data.ID, data); // edit data ke API
      Swal.fire({
        icon: "success",
        title: "Berhasil!",
        text: "Data berhasil diperbarui",
        showConfirmButton: false,
        timer: 1500,
      }).then(() => {
        setMainData(
          MainData.map((item) => {
            return item.ID === data.ID ? data : item;
          })
        );
      });
    } catch (error) {
      Swal.fire({
        icon: "error",
        title: "Gagal!",
        text: "Error saat mengubah data",
        showConfirmButton: false,
        timer: 1500,
      });
    } finally {
      onCloseFormModal();
    }
  };

  // Function untuk hapus 1 data
  const handleDelete = async (id) => {
    Swal.fire({
      title: "Apakah Anda yakin?",
      text: "Anda akan menghapus data ini!",
      icon: "warning",
      showCancelButton: true,
      confirmButtonText: "Ya, saya yakin",
      cancelButtonText: "Batal",
    }).then(async (result) => {
      if (result.isConfirmed) {
        try {
          await deleteSuratMasuk(id); // hapus data di API
          setMainData((prevData) => prevData.filter((data) => data.ID !== id));
          Swal.fire({
            icon: "info",
            title: "Berhasil!",
            text: "Data berhasil dihapus",
            showConfirmButton: false,
            timer: 1500,
          });
        } catch (error) {
          Swal.fire("Gagal!", "Error saat hapus data:", "error");
        }
      }
    });
  };

  // handle select
  const handleSelect = ({ selectedRows }) => {
    const id = selectedRows.map((data) => data.ID);
    setSelectedIds(id);
  };

  // Function untuk hapus multi select checkbox
  const handleBulkDelete = async () => {
    Swal.fire({
      title: "Apakah Anda yakin?",
      text: "Anda akan menghapus data yang dipilih!",
      icon: "warning",
      showCancelButton: true,
      confirmButtonText: "Ya, saya yakin",
      cancelButtonText: "Batal",
    }).then(async (result) => {
      if (result.isConfirmed) {
        try {
          await Promise.all(selectedIds.map((id) => deleteSuratMasuk(id))); // hapus data di API
          setMainData((prevData) =>
            prevData.filter((data) => !selectedIds.includes(data.ID))
          );
          setSelectedIds([]);
          Swal.fire({
            icon: "info",
            title: "Berhasil!",
            text: "Data berhasil dihapus",
            showConfirmButton: false,
            timer: 1500,
          });
        } catch (error) {
          Swal.fire("Gagal!", "Error saat hapus data:", "error");
        }
      }
    });
  };

  return (
    <App services={formConfig.services}>
      <div className="overflow-auto">
        {/* Table */}
        <ReusableTable
          MainData={MainData}
          formConfig={formConfig}
          handleAdd={handleAdd}
          handleEdit={handleEdit}
          handleDelete={handleDelete}
          handleSelect={handleSelect}
          selectedIds={selectedIds}
          handleBulkDelete={handleBulkDelete}
          linkExportThis=""
          linkUpdateThis=""
          importExcel=""
        />
        {/* End Table */}

        {/* ModalForm */}
        <Modal show={formModalOpen} size="xl" onClose={onCloseFormModal} popup>
          <Modal.Header />
          <Modal.Body>
            <ReusableForm
              config={formConfig}
              formData={formData}
              setFormData={setFormData}
            />
          </Modal.Body>
        </Modal>
        {/* endModalForm */}
      </div>
    </App>
  );
}
