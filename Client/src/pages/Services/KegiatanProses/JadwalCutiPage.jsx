import { useState, useEffect } from "react";
import { v4 as uuidv4 } from "uuid"; // Import UUID
import Swal from "sweetalert2";
import App from "../../../components/Layouts/App";
import { ReusableCalendar } from "../../../components/Fragments/Services/ReusableCalendar";
import {
  getCutis,
  addCuti,
  deleteCuti,
} from "../../../../API/KegiatanProses/JadwalCuti.service";

export function JadwalCutiPage() {
  const [currentEvents, setCurrentEvents] = useState([]);
  // Fetch events
  useEffect(() => {
    getCutis((data) => {
      setCurrentEvents(data.reverse());
    });
  }, []);

  // Handle date click to add new event
  const handleDateClick = async (selected) => {
    const calendarApi = selected.view.calendar;
    const { value: title } = await Swal.fire({
      title: "Masukan Event!",
      input: "text",
      inputAttributes: {
        autocapitalize: "off",
      },
      showCancelButton: true,
      confirmButtonText: "Simpan",
      showLoaderOnConfirm: true,
      preConfirm: (e) => {
        return {
          id: uuidv4(),
          title: e,
          start: selected.startStr,
          end: selected.endStr,
          allDay: selected.allDay,
        };
      },
    });
    if (title) {
      try {
        await addCuti(title);
        setCurrentEvents((prevEvents) => [...prevEvents, title]);
      } catch (error) {
        Swal.fire({
          icon: "error",
          title: "Gagal!",
          text: "Error saat menyimpan data: " + error.message,
          showConfirmButton: false,
          timer: 1500,
        });
      }
    } else {
      calendarApi.unselect();
    }
  };

  // Handle event click to delete event
  const handleEventClick = async (selected) => {
    Swal.fire({
      title: "Apakah Anda yakin?",
      text: `Anda akan menghapus data ${selected.event.title}?`,
      icon: "warning",
      showCancelButton: true,
      confirmButtonText: "Ya, saya yakin",
      cancelButtonText: "Batal",
    }).then(async (result) => {
      if (result.isConfirmed) {
        try {
          await deleteCuti(selected.event.id);
          setCurrentEvents((prevEvents) =>
            prevEvents.filter((event) => event.id !== selected.event.id)
          );
          Swal.fire({
            icon: "info",
            title: "Berhasil!",
            text: "Data berhasil dihapus",
            showConfirmButton: false,
            timer: 1500,
          });
        } catch (error) {
          // Perbaikan di sini
          Swal.fire({
            icon: "error",
            title: "Gagal!",
            text: "Error saat menghapus data: " + error.message, // Menampilkan pesan error
            showConfirmButton: false,
            timer: 1500,
          });
        }
      }
    });
  };
  return (
    <App services="Jadwal Cuti">
      <ReusableCalendar
        currentEvents={currentEvents}
        handleDateClick={handleDateClick}
        handleEventClick={handleEventClick}
      />
    </App>
  );
}
